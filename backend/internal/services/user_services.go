package services

import (
	"fmt"

	"rainbow-umbrella/internal/interfaces"
)

type userService struct {
}

func NewUserService() interfaces.IUserService {
	return &userService{}
}

func (s userService) Register() {
	fmt.Println("[userService]: register")
}
