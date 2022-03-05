package repos

import (
	"database/sql"
	"fmt"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/dao"
)

type friendshipRepo struct {
	dbClient *sql.DB
}

func NewFriendshipRepo(dbClient *sql.DB) interfaces.IFriendshipRepo {
	return &friendshipRepo{dbClient: dbClient}
}

func (r friendshipRepo) InsertOne(friendship *dao.Friendship) error {
	q := buildInsertOneFriendshipQuery(friendship)

	if _, err := r.dbClient.Exec(q.Query, q.Args...); err != nil {
		return fmt.Errorf("[friendshipRepo.InsertOne][1]: %+v", err)
	}

	return nil
}
