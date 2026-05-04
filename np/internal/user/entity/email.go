package entity

import "errors"

var (
	ErrEmailValueIssue      error = errors.New("core: [domain, user, entity, email] error: [email value issue]")
	ErrEmailCharactersIssue error = errors.New("core: [domain, user, entity, email] error: [email characters issue]")
)

type email string

func NewEmail(input string) (*email, error) {
	output := email(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self *email) validation() error {
	// TODO ввод не пустой
	// TODO проверка регуляркой на то, что это почтовый ящик
	// TODO доп проверка, на наличие символов, которые запрещены
	return nil
}

func (self *email) GetEmail() string {
	return string(*self)
}
