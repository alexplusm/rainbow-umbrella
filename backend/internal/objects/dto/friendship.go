package dto

import (
	"time"

	"rainbow-umbrella/internal/consts"
	"rainbow-umbrella/internal/objects/bo"
)

type Friendship struct {
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
