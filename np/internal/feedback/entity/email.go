package entity

import "errors"

type Email string

var (
	ErrFeedbackEmailIncorrect error = errors.New("feedback email was incorrect")
)

func NewEmail(input string) (*Email, error) {
	output := Email(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self Email) validation() error {
	// lowercase
	// TODO регулярка на соответствие email
	return nil
}
