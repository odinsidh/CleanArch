package entity

import "errors"

type SortType int

const (
	Asc SortType = iota + 1
	Desc
)

var (
	ErrInvalidSortType error = errors.New("invalid sort type")
)

func NewSortType(input int) (*SortType, error) {
	output := SortType(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil

}

func (self *SortType) validation() error {
	_, ok := avaliableSortType[*self]
	if !ok {
	}
	return nil
}

var avaliableSortType map[SortType]struct{} = map[SortType]struct{}{
	Asc:  struct{}{},
	Desc: struct{}{},
}
