package services

import (
	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/bo"
)

type friendshipService struct {
	friendshipRepo interfaces.IFriendshipRepo
}

func NewFriendshipService(friendshipRepo interfaces.IFriendshipRepo) interfaces.IFriendshipService {
	return &friendshipService{friendshipRepo: friendshipRepo}
}

func (s friendshipService) Create(value *bo.Friendship) error {
	return nil
}
