package repos

import (
	"context"
	"database/sql"
	"fmt"

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

func (r interestRepo) InsertListAndAssignToUser(ctx context.Context, userID uint64, interests []string) error {
	tx, err := r.dbClient.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[interestRepo.InsertListAndAssignToUser][1]: %w", err)
	}

	q, err := buildInsertListInterestQuery(interests)
	if err != nil {
		return fmt.Errorf("[interestRepo.InsertListAndAssignToUser][2]: %w", err)
	}

	// TODO: нужно ли вызывать запросы в рамках транзакции с контекстом?
	if _, err := tx.Exec(q.Query, q.Args...); err != nil {
		return fmt.Errorf("[interestRepo.InsertListAndAssignToUser][3]: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[interestRepo.InsertListAndAssignToUser][4]: %w", err)
	}

	//  insert new interests
	//	and get it ids
	//  assign interests ids to userId

	return nil
}
