package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/dto"
)

type friendshipController struct {
	friendshipService interfaces.IFriendshipService
}

func NewFriendshipController(friendshipService interfaces.IFriendshipService) interfaces.IFriendshipController {
	return &friendshipController{friendshipService: friendshipService}
}

func (c friendshipController) Create(w http.ResponseWriter, r *http.Request) {
	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		processError(w, http.StatusBadRequest, nil)
		log.Println(fmt.Errorf("[friendshipController.Create][1]: %+v", err))
		return
	}

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

	fmt.Printf("BODY: %+v\n", body)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
