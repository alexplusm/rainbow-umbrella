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

	q, err = buildSelectInterestsIDsQuery(interests)
	if err != nil {
		return fmt.Errorf("[interestRepo.InsertListAndAssignToUser][4]: %w", err)
	}

	rows, err := tx.Query(q.Query, q.Args...)
	if err != nil {
		return fmt.Errorf("[interestRepo.InsertListAndAssignToUser][5]: %w", err)
	}

	interestIDs := make([]uint64, 0, len(interests))
	for rows.Next() {
		var id uint64
		if err = rows.Scan(&id); err != nil {
			return fmt.Errorf("[interestRepo.InsertListAndAssignToUser][6]: %w", err)
		}
		interestIDs = append(interestIDs, id)
	}

	q, err = buildInsertListUserInterestQuery(userID, interestIDs)
	if err != nil {
		return fmt.Errorf("[interestRepo.InsertListAndAssignToUser][7]: %w", err)
	}

	if _, err = tx.Exec(q.Query, q.Args...); err != nil {
		return fmt.Errorf("[interestRepo.InsertListAndAssignToUser][8]: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("[interestRepo.InsertListAndAssignToUser][9]: %w", err)
	}

	return nil
}

func (r interestRepo) SelectListByUserID(tx *sql.Tx, ctx context.Context, userID uint64) ([]string, error) {
	var (
		rows *sql.Rows
		err  error
	)

	q := buildSelectListInterestByUserIDQuery(userID)

	if tx != nil {
		rows, err = tx.Query(q.Query, q.Args...)
	} else {
		rows, err = r.dbClient.QueryContext(ctx, q.Query, q.Args...)
	}
	if err != nil {
		return nil, fmt.Errorf("[interestRepo.SelectListByUserID][1]: %w", err)
	}

	interests := make([]string, 0, 16)

	for rows.Next() {
		var interest string

		if err = rows.Scan(&interest); err != nil {
			return nil, fmt.Errorf("[interestRepo.SelectListByUserID][2]: %w", err)
		}
		interests = append(interests, interest)
	}

	return interests, nil
}
