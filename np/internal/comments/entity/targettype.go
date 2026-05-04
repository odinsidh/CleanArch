package entity

import "errors"

type targetType string

var (
	ErrTargetTypeUnknown error = errors.New("core: [domain, comment, entity, targetType] error: [unknown target type]")
)

func NewTargetType(input string) (*targetType, error) {
	output := targetType(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self targetType) validation() error {
	// TODO проверки:
	// убедиться, что у нас существует такой targetType, который приходит при попытке создать сообщение
	// ErrTargetTypeIsNotExist
	return nil
}
