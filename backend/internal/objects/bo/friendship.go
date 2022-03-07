package bo

import (
	"time"

	"rainbow-umbrella/internal/objects/dto"
)

type Friendship struct {
	ID               uint64
	RequestingUserID uint64
	TargetingUserID  uint64
	Status           string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *Friendship) Build(value *dto.Friendship) *Friendship {
	o.RequestingUserID = value.RequestingUserID
	o.TargetingUserID = value.TargetingUserID
	o.CreatedAt = time.Now()

	return o
}
