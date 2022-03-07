package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/bo"
	"rainbow-umbrella/internal/objects/dto"
)

type friendshipController struct {
	friendshipService interfaces.IFriendshipService
}

func NewFriendshipController(friendshipService interfaces.IFriendshipService) interfaces.IFriendshipController {
	return &friendshipController{friendshipService: friendshipService}
}

func (c friendshipController) Create(w http.ResponseWriter, r *http.Request) {
	rawBody, _ := ioutil.ReadAll(r.Body)

	body := new(dto.Friendship)

	if err := json.Unmarshal(rawBody, body); err != nil {
	}

	friendship := new(bo.Friendship).Build(body)

	if err := c.friendshipService.Create(friendship); err != nil {
	}

	fmt.Printf("BODY: %+v\n", body)

	w.Write([]byte("123 aza"))
}
