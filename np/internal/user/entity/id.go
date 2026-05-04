package entity

import "errors"

var (
	ErrIdValueIssue error = errors.New("core: [domain, user, entity, id] error: [id value issue]")
)

type id int

func NewId(input int) (*id, error) {
	output := id(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self *id) validation() error {
	// TODO по идее тут валидация не нужна никакая, ибо ID мы получаем в USECASE (пока оставим это жить тут)
	return nil
}
