package repos

import (
	"context"
	"database/sql"
	"fmt"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/bo"
	"rainbow-umbrella/internal/objects/dao"
)

type userRepo struct {
	dbClient *sql.DB
}

func NewUserRepo(dbClient *sql.DB) interfaces.IUserRepo {
	return &userRepo{dbClient: dbClient}
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
