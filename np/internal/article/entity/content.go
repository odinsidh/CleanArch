package entity

import "errors"

var (
	ErrContentLenghtIssue error = errors.New("content lenght issue")
)

const (
	minimalContentLenght titleLenght = 200
	maximalContentLenght titleLenght = 1200
)

type content string

func NewContent(input string) (*content, error) {
	output := content(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self *content) validation() error {
	// TODO content lenght chech
	// TODO можно тут же проверку на маты сделать
	return nil
}
