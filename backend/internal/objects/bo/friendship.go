package bo

import (
	"time"
)

type Friendship struct {
	ID               uint64
	RequestingUserID uint64
	TargetingUserID  uint64
	Status           string

	CreatedAt time.Time
	UpdatedAt time.Time
}
