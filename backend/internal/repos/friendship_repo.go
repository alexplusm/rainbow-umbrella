package repos

import (
	"database/sql"

	"rainbow-umbrella/internal/interfaces"
)

type friendshipRepo struct {
	dbClient *sql.DB
}

func NewFriendshipRepo(dbClient *sql.DB) interfaces.IFriendshipRepo {
	return &friendshipRepo{dbClient: dbClient}
}
