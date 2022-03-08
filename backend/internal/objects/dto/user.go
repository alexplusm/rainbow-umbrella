package dto

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"rainbow-umbrella/internal/objects/bo"
)

type User struct {
	ID           uint64 `json:"id"`
	FriendshipID uint64 `json:"friendshipId,omitempty"`

	Login          string `json:"login"`
	Password       string `json:"password,omitempty"`
	HashedPassword string `json:"hashedPassword,omitempty"`

	Birthday  string `json:"birthday"` // TODO
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Gender    string `json:"gender"`
	City      string `json:"city"`

	Interests []string `json:"interests,omitempty"` // TODO: late

	//AvatarURL string // INFO: may be no need
}

func (o *User) BuildFromFormValue(form map[string][]string) (*User, error) {
	if len(form["login"]) != 0 {
		o.Login = form["login"][0]
	} else {
		return nil, fmt.Errorf("login required")
	}

	if len(form["password"]) != 0 {
		password := form["password"][0]

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("password process error: %+v", err)
		}

		o.HashedPassword = string(hashedPassword)
	} else {
		return nil, fmt.Errorf("password required")
	}

	if len(form["firstName"]) != 0 {
		o.FirstName = form["firstName"][0]
	} else {
		return nil, fmt.Errorf("first name required")
	}

	if len(form["lastName"]) != 0 {
		o.Lastname = form["lastName"][0]
	} else {
		return nil, fmt.Errorf("last name required")
	}

	if len(form["gender"]) != 0 {
		o.Gender = form["gender"][0]
	} else {
		return nil, fmt.Errorf("gender required")
	}

	if len(form["city"]) != 0 {
		o.City = form["city"][0]
	} else {
		return nil, fmt.Errorf("city required")
	}

	// TODO: late
	if len(form["interests"]) != 0 {
		o.Interests = strings.Split(form["interests"][0], ",")
	}

	return o, nil
}

func (o User) ToBO() *bo.User {
	return &bo.User{
		ID:             o.ID,
		Login:          o.Login,
		HashedPassword: o.HashedPassword,

		FirstName: o.FirstName,
		LastName:  o.Lastname,
		Birthday:  time.Now(), // TODO
		Gender:    o.Gender,
		City:      o.City,

		CreatedAt: time.Now(),

		Interests: o.Interests,
	}
}

func (o *User) FromBO(user *bo.User) *User {
	o.ID = user.ID
	o.FriendshipID = user.FriendshipID
	o.Login = user.Login

	// TODO: birthday
	o.FirstName = user.FirstName
	o.Lastname = user.LastName
	o.Gender = user.Gender
	o.City = user.City

	return o
}

// ---

type UserLoginResponse struct {
	SessionID string `json:"sessionID"`
}
