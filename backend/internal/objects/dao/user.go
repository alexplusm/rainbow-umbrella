package dao

import (
	"fmt"
	"time"

	"rainbow-umbrella/internal/objects/bo"
)

type User struct {
	ID             uint64
	FriendshipID   uint64
	Login          string
	HashedPassword string

	FirstName string
	LastName  string
	Birthday  string
	Gender    string
	City      string

	CreatedAt string

	Interests []string
}

func (o *User) FromBO(value *bo.User) *User {
	o.ID = value.ID
	o.Login = value.Login
	o.HashedPassword = value.HashedPassword

	o.FirstName = value.FirstName
	o.LastName = value.LastName
	o.Birthday = value.Birthday.Format("2006-01-02")
	o.Gender = value.Gender
	o.City = value.City

	o.CreatedAt = timeToDAO(value.CreatedAt)

	return o
}

func (o *User) ToBO() (*bo.User, error) {
	birthday, err := time.Parse("2006-01-02", o.Birthday)
	if err != nil {
		return nil, fmt.Errorf("[User.ToBO][1]: %+v", err)
	}

	createdAt, err := timeFromDAO(o.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("[User.ToBO][2]: %+v", err)
	}

	return &bo.User{
		ID:             o.ID,
		FriendshipID:   o.FriendshipID,
		Login:          o.Login,
		HashedPassword: o.HashedPassword,

		FirstName: o.FirstName,
		LastName:  o.LastName,
		Birthday:  birthday,
		Gender:    o.Gender,
		City:      o.City,

		CreatedAt: createdAt,
	}, nil
}
