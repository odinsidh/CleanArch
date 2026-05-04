package entity

import "errors"

type UserMessage string

const (
	UserMessageMaxLen int = 500
)

var (
	ErrUserMessageTooLong error = errors.New("user message was to long")
)

func NewUserMessage(input string) (*UserMessage, error) {
	output := UserMessage(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self *UserMessage) validation() error {
	if len(*self) > UserMessageMaxLen {
		return ErrUserMessageTooLong
	}

	return nil
}
