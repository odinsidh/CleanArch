package entity

import (
	"errors"
)

var (
	ErrUsernameValueIssue      error = errors.New("core: [domain, user, entity, username] error: [username value issue]")
	ErrUsernameLenghtIssue     error = errors.New("core: [domain, user, entity, username] error: [username lenght issue]")
	ErrUsernameCharactersIssue error = errors.New("core: [domain, user, entity, username] error: [username characters issue]")
)

type username string

func NewUsername(input string) (*username, error) {
	output := username(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self *username) validation() error {
	// TODO проверка, данные на входе не равны пустоте
	// TODO проверка длины юзернейма
	// TODO проверка разрешеные символов
	return nil
}
