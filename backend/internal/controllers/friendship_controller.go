package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/dto"
)

type friendshipController struct {
}

func NewFriendshipController() interfaces.IFriendshipController {
	return &friendshipController{}
}

func (c friendshipController) Create(w http.ResponseWriter, r *http.Request) {
	rawBody, _ := ioutil.ReadAll(r.Body)

	body := new(dto.Friendship)

	if err := json.Unmarshal(rawBody, body); err != nil {
	}

	fmt.Printf("BODY: %+v\n", body)

	w.Write([]byte("123 aza"))
}
