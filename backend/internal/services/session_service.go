package services

import (
	"rainbow-umbrella/internal/interfaces"
)

type sessionService struct {
	sessionRepo interfaces.ISessionRepo
}

func NewSessionService(sessionRepo interfaces.ISessionRepo) interfaces.ISessionService {
	return &sessionService{sessionRepo: sessionRepo}
}
