package entity

import "errors"

type ContentType int

const (
	Article ContentType = iota + 1
	Comment
)

var (
	ErrContentTypeNotFound error = errors.New("content type not found")
)

func NewContentType(input int) (*ContentType, error) {
	output := ContentType(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self *ContentType) validation() error {
	_, ok := avaliableContentType[*self]
	if !ok {
		return ErrContentTypeNotFound
	}

	return nil
}

var avaliableContentType map[ContentType]struct{} = map[ContentType]struct{}{
	Article: struct{}{},
	Comment: struct{}{},
}
