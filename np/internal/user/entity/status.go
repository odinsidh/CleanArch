package entity

import "errors"

var (
	ErrStatusValueIssue     error = errors.New("core: [domain, user, entity, status] error: [status value issue]")
	ErrStatusWasNotExist    error = errors.New("core: [domain, user, entity, status] error: [status was not exist]")
	ErrStatusAlreadyChanged error = errors.New("core: [domain, user, entity, status] error: [status is already changed]")
)

const (
	STATUS_ACTIVATED   string = "activated"
	STATUS_DEACTIVATED string = "deactivated"
)

type status string

func NewStatus() (*status, error) {
	output := status(STATUS_ACTIVATED)
	return &output, nil
}

func (self *status) validation() error {
	// TODO представим, что у нас тут обьявлена структура, с только существующими статусами,
	// и мы проверяем, данные на входе, попадают под существующие статусы или нет
	return nil
}

func (self *status) GetStatus() string {
	// возможно пригодится метод, пока оставляю это тут
	return string(*self)
}

func (self *status) Deactivate() error {
	var newStatus status = status(STATUS_DEACTIVATED)
	self = &newStatus
	return nil
}

func (self *status) Activate() error {
	var newStatus status = status(STATUS_ACTIVATED)
	self = &newStatus
	return nil
}
