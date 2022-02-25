package services

import (
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
	// TODO: implements
	return "new-session-ID", nil
}
