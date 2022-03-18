package dto

import (
	"time"

	"rainbow-umbrella/internal/consts"
	"rainbow-umbrella/internal/objects/bo"
)

type Friendship struct {
	ID               uint64 `json:"id"`
	RequestingUserID uint64 `json:"requestingUserId"`
	TargetingUserID  uint64 `json:"targetingUserId"`
}

func (o *Friendship) ToBO() *bo.Friendship {
	return &bo.Friendship{
		RequestingUserID: o.RequestingUserID,
		TargetingUserID:  o.TargetingUserID,
		Status:           consts.FriendshipStatusNew,
		CreatedAt:        time.Now(),
	}
}

type FriendList struct {
	Friends            []User `json:"friends"`
	Requested          []User `json:"requested"`
	WaitingForResponse []User `json:"waitingForResponse"`
}

func (o *FriendList) FromBO(value *bo.FriendList) *FriendList {
	o.Friends = []User{}
	o.Requested = []User{}
	o.WaitingForResponse = []User{}

	for _, item := range value.Friends {
		user := new(User).FromBO(&item)
		o.Friends = append(o.Friends, *user)
	}

	for _, item := range value.Requested {
		user := new(User).FromBO(&item)
		o.Requested = append(o.Requested, *user)
	}

	for _, item := range value.WaitingForResponse {
		user := new(User).FromBO(&item)
		o.WaitingForResponse = append(o.WaitingForResponse, *user)
	}

	return o
}
