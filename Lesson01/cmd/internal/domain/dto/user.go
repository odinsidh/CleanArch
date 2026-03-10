package dto

type CreateUserInput struct {
	email    string
	username string
}

type CreateUserOutput struct {
	id float64
}

type GetUserInput struct {
}

type GetUserOutput struct {
}

type DeactivateUserInput struct {
}

type DeactivateUserOutput struct {
}
