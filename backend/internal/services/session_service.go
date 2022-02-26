package services

import (
	"fmt"

	"github.com/google/uuid"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/bo"
)

type sessionService struct {
	sessionRepo interfaces.ISessionRepo
}

func NewSessionService(sessionRepo interfaces.ISessionRepo) interfaces.ISessionService {
	return &sessionService{sessionRepo: sessionRepo}
}

func (s sessionService) Create(user *bo.User) (string, error) {
	sessionID := uuid.NewString()

	if err := s.sessionRepo.InsertOne(sessionID, user.Login); err != nil {
		return "", fmt.Errorf("[sessionService.Create][1]: %+v", err)
	}

	return sessionID, nil
}

func (s sessionService) Exists(sessionID string) (bool, error) {
	exists, err := s.sessionRepo.Exists(sessionID)
	if err != nil {
		return false, fmt.Errorf("[sessionService.Exists][1]: %+v", err)
	}

	return exists, nil
}
