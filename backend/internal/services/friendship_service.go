package services

import (
	"fmt"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/bo"
	"rainbow-umbrella/internal/objects/dao"
)

type friendshipService struct {
	friendshipRepo interfaces.IFriendshipRepo
}

func NewFriendshipService(friendshipRepo interfaces.IFriendshipRepo) interfaces.IFriendshipService {
	return &friendshipService{friendshipRepo: friendshipRepo}
}

func (s friendshipService) Create(value *bo.Friendship) error {
	if err := s.friendshipRepo.InsertOne(new(dao.Friendship).FromBO(value)); err != nil {
		return fmt.Errorf("[friendshipService.Create][1]: %w", err)
	}

	return nil
}

func (s friendshipService) FriendList(user *bo.User) (*bo.FriendList, error) {
	friendList, err := s.friendshipRepo.FriendList(user.ID)
	if err != nil {
		return nil, fmt.Errorf("[friendshipService.FriendList][1]: %+v", err)
	}

	friendListBO, err := friendList.ToBO()
	if err != nil {
		return nil, fmt.Errorf("[friendshipService.FriendList][2]: %+v", err)
	}

	return friendListBO, nil
}

func (s friendshipService) UpdateStatus(id uint64, status string) error {
	if err := s.friendshipRepo.UpdateStatus(id, status); err != nil {
		return fmt.Errorf("[friendshipService.UpdateStatus][1]: %+v", err)
	}

	return nil
}
