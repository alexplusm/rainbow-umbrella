package dao

import (
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
