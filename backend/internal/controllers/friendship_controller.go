package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/dto"
)

type friendshipController struct {
	friendshipService interfaces.IFriendshipService
	userService       interfaces.IUserService
}

func NewFriendshipController(
	friendshipService interfaces.IFriendshipService,
	userService interfaces.IUserService,
) interfaces.IFriendshipController {
	return &friendshipController{friendshipService: friendshipService, userService: userService}
}

func (c friendshipController) Create(w http.ResponseWriter, r *http.Request) {
	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		processError(w, http.StatusBadRequest, nil)
		log.Println(fmt.Errorf("[friendshipController.Create][1]: %+v", err))
		return
	}

	// TODO: check userID existence
	body := new(dto.Friendship)

	if err = json.Unmarshal(rawBody, body); err != nil {
		processError(w, http.StatusBadRequest, nil)
		log.Println(fmt.Errorf("[friendshipController.Create][2]: %+v", err))
		return
	}

	if err = c.friendshipService.Create(body.ToBO()); err != nil {
		processError(w, http.StatusInternalServerError, nil)
		log.Println(fmt.Errorf("[friendshipController.Create][3]: %+v", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func (c friendshipController) List(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")

	user, err := c.userService.RetrieveByLogin(login)
	if err != nil {
		processError(w, http.StatusInternalServerError, nil)
		log.Println(fmt.Errorf("[friendshipController.List][1]: %+v", err))
		return
	}

	if user == nil {
		processError(w, http.StatusNotFound, nil)
		return
	}

	friendList, err := c.friendshipService.FriendList(user)
	if err != nil {
		processError(w, http.StatusInternalServerError, nil)
		log.Println(fmt.Errorf("[friendshipController.List][2]: %+v", err))
		return
	}

	responseBody, err := json.Marshal(new(dto.FriendList).FromBO(friendList))
	if err != nil {
		processError(w, http.StatusInternalServerError, nil)
		log.Println(fmt.Errorf("[friendshipController.List][3]: %+v", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
	w.Header().Set("Content-Type", "application/json") // TODO: why don't work // change to ADD
}
