package entity

import (
	"errors"
	"time"
)

var (
	ErrAgeIssue error = errors.New("core: [domain, user, entity, age] error: [age issue]")
)

type age time.Time

func NewAge(input time.Time) (*age, error) {
	output := age(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}
	return &output, nil
}

func (self *age) validation() error {
	// TODO проверка, что возраст пользователя > 18 лет
	return nil
}

// func defineAge()
