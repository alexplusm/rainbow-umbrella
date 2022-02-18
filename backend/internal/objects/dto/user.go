package dto

import (
	"fmt"
	"rainbow-umbrella/internal/objects/bo"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             string
	Login          string
	HashedPassword string

	FirstName string
	Lastname  string
	Gender    string
	City      string

	Interests []string // TODO: late

	//AvatarURL string // INFO: may be no need
}

func (o *User) BuildFromFormValue(form map[string][]string) (*User, error) {
	o.ID = uuid.NewString()

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
		Gender:    o.Gender,
		City:      o.City,

		Interests: o.Interests,
	}
}
