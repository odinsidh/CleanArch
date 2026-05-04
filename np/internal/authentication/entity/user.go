package entity

import "time"

type User struct {
	UserID    int
	Username  Username
	Email     Email
	Password  Password
	Timestamp time.Time
}

func NewUser(username, email, password string) (User, error) {
	validatedUsername, err := NewUsername(username)
	if err != nil {
		return User{}, err
	}

	validatedEmail, err := NewEmail(email)
	if err != nil {
		return User{}, err
	}

	validatedPassword, err := NewPassword(password)
	if err != nil {
		return User{}, err
	}

	var output User = User{
		UserID:    0,
		Username:  validatedUsername,
		Email:     validatedEmail,
		Password:  validatedPassword,
		Timestamp: time.Now(),
	}

	return output, nil
}

func (self *User) SetHashedPassword(input []byte) {
	newHashedPassword := Password(string(input))
	self.Password = newHashedPassword
}
