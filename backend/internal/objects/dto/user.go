package dto

import (
	"fmt"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"rainbow-umbrella/internal/objects/bo"
)

type User struct {
	ID             uint64 `json:"id"`
	Login          string `json:"login"`
	Password       string `json:"password,omitempty"`
	HashedPassword string `json:"hashedPassword,omitempty"`

	Birthday  string `json:"birthday,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	Lastname  string `json:"lastName,omitempty"`
	Gender    string `json:"gender,omitempty"`
	City      string `json:"city,omitempty"`

	Interests []string `json:"interests,omitempty"`

	// --- calculated

	Age int `json:"age"`
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
			return nil, fmt.Errorf("password process error: %w", err)
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

	if len(form["birthday"]) != 0 {
		o.Birthday = form["birthday"][0]
	} else {
		return nil, fmt.Errorf("birthday required")
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

	if len(form["interests"]) != 0 {
		interests := make([]string, 0, 32)
		for _, rawInterest := range strings.Split(form["interests"][0], ",") {
			interest := strings.TrimSpace(strings.ToLower(rawInterest))
			if interest != "" {
				interests = append(interests, interest)
			}
		}
		o.Interests = interests
	}

	return o, nil
}

func (o User) ToBO() *bo.User {
	user := &bo.User{
		ID:             o.ID,
		Login:          o.Login,
		HashedPassword: o.HashedPassword,

		FirstName: o.FirstName,
		LastName:  o.Lastname,
		Gender:    o.Gender,
		City:      o.City,

		Interests: o.Interests,

		CreatedAt: time.Now(),
	}

	birthday, err := time.Parse("2006/01/02", o.Birthday)
	if err != nil {
		err = fmt.Errorf("error while parsing birthday: %v: %w", o.Birthday, err)
		log.Print(err)
	} else {
		user.Birthday = birthday
	}

	return user
}

func (o *User) FromBO(user *bo.User) *User {
	o.ID = user.ID

	o.Login = user.Login

	o.FirstName = user.FirstName
	o.Lastname = user.LastName
	o.Gender = user.Gender
	o.City = user.City

	o.Interests = user.Interests

	o.Age = int(time.Since(user.Birthday) / (time.Hour * 24 * 365))

	return o
}

// ---

type UserLoginResponse struct {
	SessionID string `json:"sessionID"` // TODO: sessionId
}
