package entity

type Username string

func NewUsername(input string) (Username, error) {
	output := Username(input)
	err := output.validation()
	if err != nil {
		return "", err
	}

	return output, nil
}

func (self Username) validation() error {
	// TODO
	if self == "" {
		return ErrUsernameNotFound
	}
	// не менее длины Х
	// не более длины Y
	// не содержит запрещенные символы
	return nil
}
