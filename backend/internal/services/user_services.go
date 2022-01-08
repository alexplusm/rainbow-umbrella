package services

import (
	"fmt"
	"time"

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

func (s userService) GenerateAvatarFileName(originalName string) string {
	return fmt.Sprintf("avatar_%v_%v", time.Now().UnixNano(), originalName)
}

func (s userService) LoginExist(login string) (bool, error) {
	return true, nil
}
