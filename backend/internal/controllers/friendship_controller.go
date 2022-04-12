package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"rainbow-umbrella/internal/consts"
	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/dto"
	"rainbow-umbrella/internal/utils"
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
		log.Print(fmt.Errorf("[friendshipController.Create][1]: %w", err))
		return
	}

	body := new(dto.Friendship)

	if err = json.Unmarshal(rawBody, body); err != nil {
		processError(w, http.StatusBadRequest, nil)
		log.Print(fmt.Errorf("[friendshipController.Create][2]: %w", err))
		return
	}

	if err = c.friendshipService.Create(body.ToBO()); err != nil {
		statusCode := http.StatusInternalServerError

		if errors.Is(err, utils.AppErrorAlreadyExist) {
			statusCode = http.StatusConflict
		}

		processError(w, statusCode, nil)
		log.Print(fmt.Errorf("[friendshipController.Create][3]: %w", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte(http.StatusText(http.StatusOK))); err != nil {
		log.Print(fmt.Errorf("[friendshipController.Create][4]: %w", err))
	}
}

func (c friendshipController) List(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")

	user, err := c.userService.RetrieveByLogin(login)
	if err != nil {
		processError(w, http.StatusInternalServerError, nil)
		log.Println(fmt.Errorf("[friendshipController.List][1]: %w", err))
		return
	}
	if user == nil {
		processError(w, http.StatusNotFound, nil)
		return
	}

	friendList, err := c.friendshipService.FriendList(user)
	if err != nil {
		processError(w, http.StatusInternalServerError, nil)
		log.Println(fmt.Errorf("[friendshipController.List][2]: %w", err))
		return
	}

	responseBody, err := json.Marshal(new(dto.FriendList).FromBO(friendList))
	if err != nil {
		processError(w, http.StatusInternalServerError, nil)
		log.Print(fmt.Errorf("[friendshipController.List][3]: %w", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(responseBody); err != nil {
		log.Print(fmt.Errorf("[friendshipController.List][4]: %w", err))
	}
}

func (c friendshipController) Approve(w http.ResponseWriter, r *http.Request) {
	body := new(dto.Friendship)

	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		processError(w, http.StatusBadRequest, nil)
		log.Print(fmt.Errorf("[friendshipController.Approve][1]: %w", err))
		return
	}
	if err = json.Unmarshal(rawBody, body); err != nil {
		processError(w, http.StatusBadRequest, nil)
		log.Print(fmt.Errorf("[friendshipController.Approve][2]: %w", err))
		return
	}

	if err = c.friendshipService.UpdateStatus(body.ID, consts.FriendshipStatusFriends); err != nil {
		processError(w, http.StatusInternalServerError, nil)
		log.Print(fmt.Errorf("[friendshipController.Approve][3]: %w", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte(http.StatusText(http.StatusOK))); err != nil {
		log.Print(fmt.Errorf("[friendshipController.Approve][4]: %w", err))
	}
}

func (c friendshipController) Decline(w http.ResponseWriter, r *http.Request) {
	body := new(dto.Friendship)

	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		processError(w, http.StatusBadRequest, nil)
		log.Print(fmt.Errorf("[friendshipController.Decline][1]: %w", err))
		return
	}
	if err = json.Unmarshal(rawBody, body); err != nil {
		processError(w, http.StatusBadRequest, nil)
		log.Print(fmt.Errorf("[friendshipController.Decline][2]: %w", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte(http.StatusText(http.StatusOK))); err != nil {
		log.Print(fmt.Errorf("[friendshipController.Decline][3]: %w", err))
	}
}
