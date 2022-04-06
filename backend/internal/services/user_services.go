package services

import (
	"context"
	"fmt"
	"time"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/bo"
	"rainbow-umbrella/internal/objects/dao"
)

type userService struct {
	userRepo interfaces.IUserRepo

	interestService interfaces.IInterestService
}

func NewUserService(
	userRepo interfaces.IUserRepo,
	interestService interfaces.IInterestService,
) interfaces.IUserService {
	return &userService{userRepo: userRepo}
}

func (s userService) Register(user *bo.User) error {
	fmt.Printf("[userService]: register: %+v\n", user)

	userID, err := s.userRepo.InsertOne(context.TODO(), new(dao.User).FromBO(user))
	if err != nil {
		return fmt.Errorf("[userService.Register][1]: %+v", err)
	}

	if err := s.interestService.CreateListForUser(context.TODO(), userID, user.Interests); err != nil {
		return fmt.Errorf("[userService.Register][1]: %w", err)
	}

	return nil
}

func (s userService) LoginExist(login string) (bool, error) {
	user, err := s.RetrieveByLogin(login)
	if err != nil {
		return false, fmt.Errorf("[userService.LoginExist][1]: %+v", err)
	}

	return user != nil, nil
}

func (s userService) RetrieveByLogin(login string) (*bo.User, error) {
	list, err := s.List(&bo.UserFilter{ByLogin: login})
	if err != nil {
		return nil, fmt.Errorf("[userService.RetrieveByLogin][1]: %+v", err)
	}

	if len(list) != 1 {
		return nil, nil
	}

	user := list[0]

	return &user, nil
}

func (s userService) List(filter *bo.UserFilter) ([]bo.User, error) {
	listDAO, err := s.userRepo.List(filter)
	if err != nil {
		return nil, fmt.Errorf("[userService.List][1]: %+v", err)
	}

	listBO := make([]bo.User, 0, len(listDAO))

	for _, userDAO := range listDAO {
		userBO, err := userDAO.ToBO()
		if err != nil {
			return nil, fmt.Errorf("[userService.List][2]: %+v", err)
		}

		listBO = append(listBO, *userBO)
	}

	return listBO, nil
}

func (s userService) GenerateAvatarFileName(originalName string) string {
	return fmt.Sprintf("avatar_%v_%v", time.Now().UnixNano(), originalName)
}
