package entity

import "errors"

type SortType int

const (
	Asc SortType = iota + 1
	Desc
)

var (
	ErrSortTypeIssue error = errors.New("requested sortype is not exist")
)

var sortTypeAvaliable map[SortType]bool = map[SortType]bool{
	Asc:  true,
	Desc: true,
}

func NewSortType(input int) (*SortType, error) {
	output := SortType(input)
	err := output.validation()
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (self *SortType) validation() error {
	_, ok := sortTypeAvaliable[*self]
	if !ok {
		return ErrSortTypeIssue
	}
	return nil
}
