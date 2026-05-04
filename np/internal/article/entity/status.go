package entity

type status int

const (
	Draft status = iota + 1
	Published
)

func NewStatus() (*status, error) {
	output := status(Draft)
	return &output, nil
}

func (self *status) Publish() {
	*self = Published
}
