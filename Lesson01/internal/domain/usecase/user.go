package usecase

import (
	"context"
	"enviromentKey/CleanArchL/Lesson01/cmd/internal/domain/dto"
	"errors"
)

type UserStorage interface {
	CreateUser(ctx context.Context, input dto.CreateUserInput) (dto.CreateUserOutput, error)
	GetUser(ctx context.Context, input dto.GetUserInput) (dto.GetUserOutput, error)
	DeactivateUser(ctx context.Context, input dto.DeactivateUserInput) (dto.DeactivateUserOutput, error)
}

var (
	ErrEmailAlreadyExist = errors.New("where: usecase [User] error: [Email Already Exist]")
)

type User struct {
	us UserStorage
}

func NewUserUsecase(userStorage UserStorage) *User {
	return &User{us: userStorage}
}

func (u *User) CreateUser()
