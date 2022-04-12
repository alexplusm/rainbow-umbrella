package services

import (
	"context"
	"fmt"
	"time"

	"rainbow-umbrella/internal/consts"
	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/bo"
	"rainbow-umbrella/internal/objects/dao"
)

type userService struct {
	userRepo interfaces.IUserRepo

	interestService   interfaces.IInterestService
	friendshipService interfaces.IFriendshipService
}

func NewUserService(
	userRepo interfaces.IUserRepo,
	interestService interfaces.IInterestService,
	friendshipService interfaces.IFriendshipService,
) interfaces.IUserService {
	return &userService{
		userRepo:          userRepo,
		interestService:   interestService,
		friendshipService: friendshipService,
	}
}

func (s userService) Register(user *bo.User) error {
	userID, err := s.userRepo.InsertOne(context.TODO(), new(dao.User).FromBO(user))
	if err != nil {
		return fmt.Errorf("[userService.Register][1]: %+v", err)
	}

	if len(user.Interests) == 0 {
		return nil
	}

	if err = s.interestService.CreateListForUser(context.TODO(), userID, user.Interests); err != nil {
		return fmt.Errorf("[userService.Register][2]: %w", err)
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
	user, err := s.userRepo.SelectOne(context.TODO(), login)
	if err != nil {
		return nil, fmt.Errorf("[userService.RetrieveByLogin][1]: %w", err)
	}

	if user == nil {
		return nil, nil
	}

	userBO, err := user.ToBO()
	if err != nil {
		return nil, fmt.Errorf("[userService.RetrieveByLogin][2]: %w", err)
	}

	return userBO, nil
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

func (s userService) GetUsersFriendshipStatus(login1, login2 string) (string, error) {
	users, err := s.listCommonInfo(&bo.UserFilter{ByLogins: []string{login1, login2}})
	if err != nil {
		return "", fmt.Errorf("[userService.GetUsersFriendshipStatus][1]: %w", err)
	}
	if len(users) != 2 {
		return consts.FriendshipStatusNotFriends, nil
	}

	friendship, err := s.friendshipService.RetrieveByUsersID(users[0].ID, users[1].ID)
	if err != nil {
		return "", fmt.Errorf("[userService.GetUsersFriendshipStatus][2]: %w", err)
	}
	if friendship == nil {
		return consts.FriendshipStatusNotFriends, nil
	}

	return friendship.Status, nil
}

func (s userService) listCommonInfo(filter *bo.UserFilter) ([]bo.UserCommonInfo, error) {
	users, err := s.userRepo.ListCommonInfo(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("[userService.listCommonInfo][1]: %w", err)
	}

	usersBO := make([]bo.UserCommonInfo, 0, len(users))
	for _, user := range users {
		usersBO = append(usersBO, *user.ToBO())
	}

	return usersBO, nil
}
