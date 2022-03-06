package controllers

import (
	"rainbow-umbrella/internal/interfaces"
)

type friendshipController struct {
}

func NewFriendshipController() interfaces.IFriendshipController {
	return &friendshipController{}
}
