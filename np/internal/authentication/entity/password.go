package entity

type Password string

func NewPassword(input string) (Password, error) {
	output := Password(input)
	err := output.validate()
	if err != nil {
		return "", err
	}

	return output, nil
}

func (self Password) validate() error {
	// TODO
	// avaliable characters
	// min lenght
	// max lenght
	// but if it was cached, we dont care
	return nil
}
