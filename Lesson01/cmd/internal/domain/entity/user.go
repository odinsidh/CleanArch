package entity

import (
	"errors"
	"time"
)

type User struct {
	Id       float64
	Email    string
	Username string
	Status   status
	Created  time.Time
}

type status int

const (
	active status = iota + 1
	deactivated
)

var (
	ErrEmailIsMissing    = errors.New("where: entity [User] error: [Email Is Missing]")
	ErrUsernameIsMissing = errors.New("where: entity [User] error: [Username Is Missing]")
)

func (u *User) CreateUser(Email string, Username string) (*User, error) {
	if err := u.validation(Email, Username); err != nil {
		return nil, err
	}
	output := &User{
		Email:    Email,
		Username: Username,
		Status:   active,
		Created:  time.Now(),
	}
	return output, nil
}

func (u *User) validation(Email string, Username string) error {
	var container []struct {
		Name  string
		Value string
		Err   error
	}
	container = append(container,
		struct {
			Name  string
			Value string
			Err   error
		}{Name: "Email Cheking", Value: Email, Err: ErrEmailIsMissing})
	container = append(container,
		struct {
			Name  string
			Value string
			Err   error
		}{Name: "Username Cheking", Value: Username, Err: ErrEmailIsMissing})

	for _, concreteContainer := range container {
		if concreteContainer.Value == "" {
			return concreteContainer.Err
		}
	}

	return nil
}
