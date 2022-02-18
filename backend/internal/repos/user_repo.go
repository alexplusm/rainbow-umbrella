package repos

import (
	"database/sql"
	"fmt"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/dao"
)

type userRepo struct {
	dbClient *sql.DB
}

func NewUserRepo(dbClient *sql.DB) interfaces.IUserRepo {
	return &userRepo{dbClient: dbClient}
}

func (r userRepo) InsertOne(item *dao.User) error {
	q := buildInsertOneUser(item)

	if _, err := r.dbClient.Exec(q.Query, q.Args); err != nil {
		return fmt.Errorf("[userRepo.InsertOne][1]: %+v", err)
	}

	return nil
}
