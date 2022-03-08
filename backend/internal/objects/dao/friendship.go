package dao

import (
	"fmt"

	"rainbow-umbrella/internal/objects/bo"
)

type Friendship struct {
	ID               uint64
	RequestingUserID uint64
	TargetingUserID  uint64
	Status           string

	CreatedAt string
	UpdatedAt string
}

func (o *Friendship) FromBO(value *bo.Friendship) *Friendship {
	o.RequestingUserID = value.RequestingUserID
	o.TargetingUserID = value.TargetingUserID
	o.Status = value.Status
	o.CreatedAt = timeToDAO(value.CreatedAt)

	return o
}

type FriendList struct {
	Friends            []User
	Requested          []User
	WaitingForResponse []User
}

func (o FriendList) ToBO() (*bo.FriendList, error) {
	friendList := new(bo.FriendList)

	for _, item := range o.Friends {
		user, err := item.ToBO()
		if err != nil {
			return nil, fmt.Errorf("[FriendList.ToBO][1]: %+v", err)
		}
		friendList.Friends = append(friendList.Friends, *user)
	}

	for _, item := range o.Requested {
		user, err := item.ToBO()
		if err != nil {
			return nil, fmt.Errorf("[FriendList.ToBO][1]: %+v", err)
		}
		friendList.Requested = append(friendList.Requested, *user)
	}

	for _, item := range o.WaitingForResponse {
		user, err := item.ToBO()
		if err != nil {
			return nil, fmt.Errorf("[FriendList.ToBO][1]: %+v", err)
		}
		friendList.WaitingForResponse = append(friendList.WaitingForResponse, *user)
	}

	return friendList, nil
}
