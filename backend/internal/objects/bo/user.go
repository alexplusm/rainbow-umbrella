package bo

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uint64
	FriendshipID   uint64
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

func (o *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(o.HashedPassword), []byte(password))

	return err == nil
}

type UserFilter struct {
	ByLogin      string
	ExcludeLogin string
}
