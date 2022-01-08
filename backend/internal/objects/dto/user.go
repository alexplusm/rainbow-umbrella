package dto

import (
	"fmt"
	"strings"
)

type User struct {
	ID        string
	Login     string
	FirstName string
	Lastname  string
	Sex       string
	City      string
	Interests []string
}

func (o *User) BuildFromFormValue(form map[string][]string) (*User, error) {
	if len(form["login"]) != 0 {
		o.Login = form["login"][0]
	} else {
		return nil, fmt.Errorf("login required")
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

	if len(form["sex"]) != 0 {
		o.Sex = form["sex"][0]
	} else {
		return nil, fmt.Errorf("sex required")
	}

	if len(form["city"]) != 0 {
		o.City = form["city"][0]
	} else {
		return nil, fmt.Errorf("city required")
	}

	if len(form["interests"]) != 0 {
		o.Interests = strings.Split(form["interests"][0], ",")
	}

	return o, nil
}
