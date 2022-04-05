package repos

import (
	"context"
	"database/sql"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/dao"
)

type interestRepo struct {
	dbClient *sql.DB
}

func NewInterestRepo(dbClient *sql.DB) interfaces.IInterestRepo {
	return &interestRepo{dbClient: dbClient}
}

func (r interestRepo) InsertOne(ctx context.Context, value *dao.Interest) error {
	q := buildInsertOneInterestQuery(value)

	if _, err := r.dbClient.ExecContext(ctx, q.Query, q.Args...); err != nil {
		return err
	}

	return nil
}
