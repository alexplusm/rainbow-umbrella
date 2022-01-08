package controllers

import (
	"fmt"
	"log"
	"net/http"

	"rainbow-umbrella/internal/interfaces"
)

type userController struct {
	userService interfaces.IUserService
}

func NewUserController(userService interfaces.IUserService) interfaces.IUserController {
	return &userController{userService: userService}
}

func (c userController) Register(w http.ResponseWriter, r *http.Request) {

	//m , err := http.Request.MultipartReader()
	//c.Register()

	if err := r.ParseMultipartForm(1000); err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.Form)

	w.Write([]byte("kekes"))
}
