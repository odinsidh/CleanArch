package entity

import (
	"errors"
)

type RequestType string

var (
	ErrRequestTypeNotExist error = errors.New("feedback request type not exist")
)

const (
	Bag      RequestType = "Баг"
	Idea     RequestType = "Идея"
	Report   RequestType = "Жалоба"
	Question RequestType = "Вопрос"
)

var registredRequestType map[RequestType]struct{} = map[RequestType]struct{}{
	Bag:      struct{}{},
	Idea:     struct{}{},
	Report:   struct{}{},
	Question: struct{}{},
}

func NewRequestType(input string) (*RequestType, error) {
	output := RequestType(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}
	return &output, nil
}

func (self RequestType) validation() error {
	_, ok := registredRequestType[self]
	if !ok {
		return ErrRequestTypeNotExist
	}
	return nil
}
