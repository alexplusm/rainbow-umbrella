package services

import (
	"rainbow-umbrella/internal/interfaces"
)

type userService struct {
}

func NewUserService() interfaces.IUserService {
	return &userService{}
}

func (s userService) Register() {}
