package bo

import (
	"log"
	"strconv"
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

	Interests []string // TODO: late

	CreatedAt time.Time
}

func (o *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(o.HashedPassword), []byte(password))

	return err == nil
}

type UserFilter struct {
	ByLogin      string
	ExcludeLogin string

	Search string
	Limit  int
	Offset int
}

func (o *UserFilter) Build() *UserFilter {
	o.Limit = 100
	return o
}

func (o *UserFilter) SetLimitAndOffset(limit, offset string) *UserFilter {
	limitV, err := strconv.Atoi(limit)
	if err != nil {
		log.Printf("[UserFilter.SetLimitAndOffset][1]: %v", err)
		o.Limit = 100
	} else {
		o.Limit = limitV
	}

	offsetV, err := strconv.Atoi(offset)
	if err != nil {
		log.Printf("[UserFilter.SetLimitAndOffset][1]: %v", err)
		o.Offset = 0
	} else {
		o.Offset = offsetV
	}

	return o
}

func (o *UserFilter) SetSearch(search string) *UserFilter {
	o.Search = search
	return o
}
