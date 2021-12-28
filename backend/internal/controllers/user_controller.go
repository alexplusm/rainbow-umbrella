package controllers

import (
	"rainbow-umbrella/internal/interfaces"
)

type userController struct {
	userService interfaces.IUserService
}

func NewUserController(userService interfaces.IUserService) interfaces.IUserController {
	return &userController{userService: userService}
}
