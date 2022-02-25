package bo

import (
	"time"
)

type User struct {
	ID             int64
	Login          string
	HashedPassword string

	FirstName string
	LastName  string
	Birthday  time.Time
	Gender    string
	City      string

	CreatedAt time.Time

	Interests []string // TODO: late
}

type UserFilter struct {
	ByLogin string
}
