package entity

import "errors"

type Status int

const (
	New Status = iota + 1
	InWork
	Banned
	Disagree
)

var (
	ErrStatusNotFound error = errors.New("status not found")
)

func NewStatus(input int) (*Status, error) {
	output := Status(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self *Status) validation() error {
	_, ok := avaliableStatus[*self]
	if !ok {
		return ErrStatusNotFound
	}

	return nil
}

func GetStatusInWork() int {
	return int(InWork)
}

var avaliableStatus map[Status]struct{} = map[Status]struct{}{
	New:      struct{}{},
	InWork:   struct{}{},
	Banned:   struct{}{},
	Disagree: struct{}{},
}
