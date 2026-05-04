package entity

import "errors"

type Message string

const (
	MessageMinLenght int = 10
	MessageMaxLenght int = 2000
)

// Вот к именованию ошибок есть вопросы, оставить как есть или лучше ErrFeedbackMessageToSmall ?
var (
	ErrMessageLenghtIssueTooSmall error = errors.New("feedback message too small")
	ErrMessageLenghtIssueTooBig   error = errors.New("feedback message too big")
)

func NewMessage(input string) (*Message, error) {
	output := Message(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self Message) validation() error {
	if len(self) > MessageMinLenght {
		return ErrMessageLenghtIssueTooSmall
	}
	if len(self) > MessageMaxLenght {
		return ErrMessageLenghtIssueTooBig
	}

	return nil
}
