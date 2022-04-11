package bo

import (
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserCommonInfo struct {
	ID        uint64
	Login     string
	FirstName string
	LastName  string
}

type User struct {
	ID             uint64 // TODO: common
	FriendshipID   uint64
	Login          string // TODO: common
	HashedPassword string

	FirstName string // TODO: common
	LastName  string // TODO: common
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

	ByLogins []string

	Search string
	Limit  int
	Offset int
}

//	TODO: wtf?
func (o UserFilter) ByLoginsToInterface() []interface{} {
	result := make([]interface{}, 0, len(o.ByLogins))
	for _, login := range o.ByLogins {
		result = append(result, login)
	}
	return result
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
