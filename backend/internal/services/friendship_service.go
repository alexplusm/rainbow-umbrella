package services

import (
	"rainbow-umbrella/internal/interfaces"
)

type friendshipService struct {
	friendshipRepo interfaces.IFriendshipRepo
}

func NewFriendshipService(friendshipRepo interfaces.IFriendshipRepo) interfaces.IFriendshipService {
	return &friendshipService{friendshipRepo: friendshipRepo}
}
