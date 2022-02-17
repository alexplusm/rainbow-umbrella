package services

import (
	"fmt"
	"time"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/bo"
)

type userService struct {
}

func NewUserService() interfaces.IUserService {
	return &userService{}
}

func (s userService) Register(user *bo.User) error {
	fmt.Printf("[userService]: register: %+v\n", user)
	return nil
}

func (s userService) GenerateAvatarFileName(originalName string) string {
	return fmt.Sprintf("avatar_%v_%v", time.Now().UnixNano(), originalName)
}

func (s userService) LoginExist(login string) (bool, error) {
	return true, nil
}
