package entity

import "errors"

type Reason int

const (
	Spam Reason = iota + 1
	BadLanguage
	NSFW
	Another
)

var (
	ErrReasonNotFound error = errors.New("reason not found")
)

func NewReason(input int) (*Reason, error) {
	output := Reason(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self *Reason) validation() error {
	_, ok := avaliableReason[*self]
	if !ok {
		return ErrReasonNotFound
	}

	return nil
}

var avaliableReason map[Reason]struct{} = map[Reason]struct{}{
	Spam:        struct{}{},
	BadLanguage: struct{}{},
	NSFW:        struct{}{},
	Another:     struct{}{},
}
