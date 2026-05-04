package entity

type Login struct {
	Username Username
	Password Password
}

func NewLogin(username, password string) (Login, error) {
	validatedUsername, err := NewUsername(username)
	if err != nil {
		return Login{}, err
	}

	validatedPassword, err := NewPassword(password)
	if err != nil {
		return Login{}, err
	}

	var output Login = Login{
		Username: validatedUsername,
		Password: validatedPassword,
	}

	return output, nil
}
