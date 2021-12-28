package controllers

import (
	"rainbow-umbrella/internal/interfaces"
)

type userController struct {
}

func NewUserController() interfaces.IUserController {
	return &userController{}
}
