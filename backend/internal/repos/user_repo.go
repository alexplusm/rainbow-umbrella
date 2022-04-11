package repos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/bo"
	"rainbow-umbrella/internal/objects/dao"
)

type userRepo struct {
	dbClient *sql.DB

	interestRepo interfaces.IInterestRepo
}

func NewUserRepo(dbClient *sql.DB, interestRepo interfaces.IInterestRepo) interfaces.IUserRepo {
	return &userRepo{dbClient: dbClient, interestRepo: interestRepo}
}

func (r userRepo) InsertOne(ctx context.Context, item *dao.User) (uint64, error) {
	q := buildInsertOneUser(item)

	result, err := r.dbClient.ExecContext(ctx, q.Query, q.Args...)
	if err != nil {
		return 0, fmt.Errorf("[userRepo.InsertOne][1]: %w", err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("[userRepo.InsertOne][2]: %w", err)
	}

	return uint64(userID), nil
}

func (r userRepo) SelectOne(ctx context.Context, login string) (*dao.User, error) {
	q := buildSelectOneUserQuery(login)

	// TODO: нужно ли делать транзакцию для нескольких селектов? - da, только нужно выбрать уровень изоляции
	tx, err := r.dbClient.BeginTx(ctx, nil)
	if err != nil {
		// TODO: rollback?
		return nil, fmt.Errorf("[userRepo.SelectOne][1]")
	}

	row := tx.QueryRow(q.Query, q.Args...)

	if err := row.Err(); err != nil {
		// TODO: rollback?
		return nil, fmt.Errorf("[userRepo.SelectOne][2]")
	}

	user := new(dao.User)
	err = row.Scan(
		&user.ID, &user.Login, &user.HashedPassword,
		&user.FirstName, &user.LastName, &user.Birthday, &user.Gender, &user.City)
	if err != nil {
		// TODO: rollback?
		if errors.Is(err, sql.ErrNoRows) {
			if err := tx.Commit(); err != nil {
				log.Print(fmt.Errorf("[userRepo.SelectOne][3.1]: %w", err))
			}
			return nil, nil
		}

		return nil, fmt.Errorf("[userRepo.SelectOne][3.2]")
	}

	interests, err := r.interestRepo.SelectListByUserID(tx, ctx, user.ID)
	if err != nil {
		// TODO: rollback?
		return nil, fmt.Errorf("[userRepo.SelectOne][4]")
	}

	if err := tx.Commit(); err != nil {
		log.Println(fmt.Errorf("[userRepo.SelectOne][5]: %w", err))
	}

	user.Interests = interests

	return user, nil
}

//	 TODO: refactor
func (r userRepo) List(filter *bo.UserFilter) ([]dao.User, error) {
	q, err := buildListUserQuery(filter)
	if err != nil {
		return nil, fmt.Errorf("[userRepo.List][1]: %+v", err)
	}

	fmt.Printf("[userRepo.List]: query: %+v\n", q)

	rows, err := r.dbClient.Query(q.Query, q.Args...)
	if err != nil {
		return nil, fmt.Errorf("[userRepo.List][2]: %+v", err)
	}

	list := make([]dao.User, 0, 128)

	for rows.Next() {
		user := new(dao.User)

		err := rows.Scan(
			&user.ID, &user.Login, &user.HashedPassword,
			&user.FirstName, &user.LastName, &user.Birthday, &user.Gender, &user.City,
			&user.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("[userRepo.List][3]: %+v", err)
		}

		list = append(list, *user)
	}

	return list, nil
}

func (r userRepo) ListCommonInfo(ctx context.Context, filter *bo.UserFilter) ([]dao.UserCommonInfo, error) {
	q, err := buildListUserCommonInfoQuery(filter)
	if err != nil {
		return nil, fmt.Errorf("[userRepo.ListCommonInfo][1]: %w", err)
	}

	users := make([]dao.UserCommonInfo, 0, 64)

	fmt.Println("Query: ", q.Query, "\n\n", q.Args)

	rows, err := r.dbClient.QueryContext(ctx, q.Query, q.Args...)
	if err != nil {
		return nil, fmt.Errorf("[userRepo.ListCommonInfo][2]: %w", err)
	}

	for rows.Next() {
		user := new(dao.UserCommonInfo)

		err := rows.Scan(&user.ID, &user.Login, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, fmt.Errorf("[userRepo.ListCommonInfo][3]: %w", err)
		}
		users = append(users, *user)
	}

	return users, nil
}
