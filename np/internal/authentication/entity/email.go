package entity

type Email string

func NewEmail(input string) (Email, error) {
	output := Email(input)
	err := output.validation()
	if err != nil {
		return "", err
	}

	return output, nil
}

func (self Email) validation() error {
	// TODO
	// only from trusted domain
	// avaliable characters
	// valid email regex
	return nil
}
