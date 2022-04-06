package services

import (
	"context"
	"fmt"

	"rainbow-umbrella/internal/interfaces"
)

type interestService struct {
	interestRepo interfaces.IInterestRepo
}

func NewInterestService(interestRepo interfaces.IInterestRepo) interfaces.IInterestService {
	return &interestService{interestRepo: interestRepo}
}

func (s interestService) CreateListForUser(ctx context.Context, userID uint64, interests []string) error {
	if err := s.interestRepo.InsertListAndAssignToUser(ctx, userID, interests); err != nil {
		return fmt.Errorf("[interestService.CreateListForUser][1]: %w", err)
	}

	return nil
}
