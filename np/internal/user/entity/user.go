package entity

import (
	"newsportal/internal/user/dto"
	"time"
)

type User struct {
	id       id
	username username
	email    email
	password password
	status   status
	age      age
	created  time.Time
}

func NewUser(username string, email string, password string, age time.Time) (*User, error) {
	userUsername, err := NewUsername(username)
	if err != nil {
		return nil, err
	}
	userEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}
	userPassword, err := NewPassword(password)
	if err != nil {
		return nil, err
	}
	userStatus, err := NewStatus()
	if err != nil {
		return nil, err
	}
	userAge, err := NewAge(age)
	if err != nil {
		return nil, err
	}
	userCreated := time.Now()

	output := &User{
		username: *userUsername,
		email:    *userEmail,
		password: *userPassword,
		status:   *userStatus,
		age:      *userAge,
		created:  userCreated,
	}

	return output, nil
}

func LoadUser(dto dto.User) (*User, error) {
	return nil, nil
}

func (self *User) DeactivateUser() error {
	currentStatus := self.status.GetStatus()
	if currentStatus == STATUS_DEACTIVATED {
		return ErrStatusAlreadyChanged
	}

	err := self.status.Deactivate()
	if err != nil {
		return err
	}

	return nil
}

func (self *User) CanBeDeactivateByUser(userID int, globalPermission bool) bool {
	if globalPermission || self.id == id(userID) {
		return true
	}
	return false
}

func (self *User) GetEmail() string {
	return self.email.GetEmail()
}
