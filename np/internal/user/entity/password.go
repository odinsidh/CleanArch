package entity

import "errors"

var (
	ErrPasswordValueIssue      error = errors.New("core: [domain, user, entity, password] error: [password value issue]")
	ErrPasswordLenghtIssue     error = errors.New("core: [domain, user, entity, password] error: [password lenght issue]")
	ErrPasswordCharactersIssue error = errors.New("core: [domain, user, entity, password] error: [password characters issue]")
)

type password string

func NewPassword(input string) (*password, error) {
	output := password(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self *password) validation() error {
	// TODO убедиться, что пароль соответствует некой длине
	// TODO убедиться, что пароль не содержить запрещенных символов
	// TODO убедиться, что пароль не пустой
	return nil
}
