package repos

import (
	"context"
	"database/sql"
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

func (r userRepo) RetrieveOne(ctx context.Context, login string) (*dao.User, error) {
	q := buildRetrieveOneUserQuery(login)

	// TODO: нужно ли делать транзакцию для нескольких селектов? - da, только нужно выбрать уровень изоляции

	tx, err := r.dbClient.BeginTx(ctx, nil)
	if err != nil {
		// TODO: rollback?
		return nil, fmt.Errorf("[userRepo.RetrieveOne][1]")
	}

	row := tx.QueryRow(q.Query, q.Args...)

	if err := row.Err(); err != nil {
		// TODO: rollback?
		return nil, fmt.Errorf("[userRepo.RetrieveOne][2]")
	}

	user := new(dao.User)

	err = row.Scan(
		&user.ID, &user.Login,
		&user.FirstName, &user.LastName, &user.Birthday, &user.Gender, &user.City)

	if err != nil {
		// TODO: rollback?
		return nil, fmt.Errorf("[userRepo.RetrieveOne][3]")
	}

	interests, err := r.interestRepo.SelectListByUserID(tx, ctx, user.ID)
	if err != nil {
		// TODO: rollback?
		return nil, fmt.Errorf("[userRepo.RetrieveOne][4]")
	}

	if err := tx.Commit(); err != nil {
		log.Println(fmt.Errorf("[userRepo.RetrieveOne][5]: %w", err))
	}

	user.Interests = interests

	return user, nil
}

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
