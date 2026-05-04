package entity

import "errors"

type titleLenght int

const (
	minimalTitleLenght titleLenght = 200
	maximalTitleLenght titleLenght = 1200
)

var (
	ErrTitleLenghtIssue error = errors.New("title lenght issue")
)

type title string

func NewTitle(input string) (*title, error) {
	output := title(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, err
}

func (self *title) validation() error {
	// TODO content lenght chech
	// TODO можно тут же проверку на маты сделать
	return nil
}
