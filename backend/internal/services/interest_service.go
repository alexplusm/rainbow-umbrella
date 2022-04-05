package services

import "rainbow-umbrella/internal/interfaces"

type interestService struct {
	interestRepo interfaces.IInterestRepo
}

func NewInterestService(interestRepo interfaces.IInterestRepo) interfaces.IInterestService {
	return &interestService{interestRepo: interestRepo}
}
