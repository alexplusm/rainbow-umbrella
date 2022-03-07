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
		return fmt.Errorf("[friendshipService.Create][1]: %+v", err)
	}

	return nil
}

func (s friendshipService) FriendList(user *bo.User) (*bo.FriendList, error) {
	friendList := &bo.FriendList{}
	friendList.Friends = make([]bo.User, 0)

	s.friendshipRepo.FriendList(user.ID)

	return friendList, nil
}
