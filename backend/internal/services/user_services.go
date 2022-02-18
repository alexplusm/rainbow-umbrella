package services

import (
	"fmt"
	"time"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/bo"
	"rainbow-umbrella/internal/objects/dao"
)

type userService struct {
	userRepo interfaces.IUserRepo
}

func NewUserService(userRepo interfaces.IUserRepo) interfaces.IUserService {
	return &userService{userRepo: userRepo}
}

func (s userService) Register(user *bo.User) error {
	fmt.Printf("[userService]: register: %+v\n", user)

	if err := s.userRepo.InsertOne(new(dao.User).FromBO(user)); err != nil {
		return fmt.Errorf("[userService.Register][1]: %+v", err)
	}

	return nil
}

func (s userService) GenerateAvatarFileName(originalName string) string {
	return fmt.Sprintf("avatar_%v_%v", time.Now().UnixNano(), originalName)
}

func (s userService) LoginExist(login string) (bool, error) {
	return true, nil
}
