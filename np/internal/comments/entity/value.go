package entity

import "errors"

type value string

var (
	ErrCommentValueIssue        error = errors.New("core: [domain, comment, entity, value] error: [value is empty]")
	ErrCommentValueMinLenght    error = errors.New("core: [domain, comment, entity, value] error: [value minimal lenght issue]")
	ErrCommentValueMaxLenght    error = errors.New("core: [domain, comment, entity, value] error: [value minimal lenght issue]")
	ErrCommentValueContentIssue error = errors.New("core: [domain, comment, entity, value] error: [value contain content issue]")
)

func NewValue(input string) (*value, error) {
	output := value(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self value) validation() error {
	// TODO проверки:
	// не равно пустоте ErrCommentValueIssue
	// есть минимальная длина ErrCommentValueMinLenght
	// есть максимальная длина ErrCommentValueMaxLenght
	// не содержит матов ErrCommentValueContentIssue
	return nil
}
