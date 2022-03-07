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
