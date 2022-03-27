package repos

import (
	"database/sql"

	"rainbow-umbrella/internal/interfaces"
)

type interestRepo struct {
	dbClient *sql.DB
}

func NewInterestRepo(dbClient *sql.DB) interfaces.IInterestRepo {
	return interestRepo{dbClient: dbClient}
}

func (r interestRepo) InsertOne() {
	buildInsertOneInterestQuery()
}
