package repos

import (
	"database/sql"

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

	//r.dbClient.QueryRow()

	return nil
}
