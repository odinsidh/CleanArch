package usecase

import (
	"context"
	dAuthN "newsportal/internal/authentication/dto"
	eAuthN "newsportal/internal/authentication/entity"
	"time"
)

// =====================================================================
//  Another interfaces
// =====================================================================

type Notification interface {
	SendMessage(ctx context.Context, actionType string, requestBody map[string]string) error
}

// =====================================================================
//  Query UseCase && Repository
// =====================================================================

type QueryUseCase interface {
	GetUserIDByUsername(ctx context.Context, username string) (userID int, err error)
	DoesUsernameAvaliable(ctx context.Context, username string) error
}

type QueryRepository interface {
	UsernameIsExist(ctx context.Context, username string) (bool, error)
	UserID(ctx context.Context, username string) (userID int, err error)
	UserIDAndPassword(ctx context.Context, username string) (UserIDAndPassword, error)
}

type QueryRepositoryCache interface {
	UsernameIsExist(ctx context.Context, username string) (bool, error)
}

// =====================================================================
//  Command UseCase && Repository
// =====================================================================

type CommandUseCase interface {
	Register(ctx context.Context, request dAuthN.CreateUser) error
	EmailConfirm(ctx context.Context, request dAuthN.EmailConfirm) error

	LogIn(ctx context.Context, request dAuthN.Login) (dAuthN.Cookie, error)
	LogOut(ctx context.Context, request dAuthN.Logout) error

	ResetPasswordRequest(ctx context.Context, request dAuthN.ResetRequire) error
	ResetPasswordConfirm(ctx context.Context, request dAuthN.ResetConfirm) error
}

type CommandRepository interface {
	Register(ctx context.Context, Username, Email, Password string, Timestamp time.Time) error
	ChangePassword(ctx context.Context, userID int, newPassword string) error
}

type CommandRepositoryCache interface {
	RegisterAdd(ctx context.Context, registerKey string, request eAuthN.User) error
	RegisterGet(ctx context.Context, registerKey string) (response eAuthN.User, err error)
	RegisterDelete(ctx context.Context, registerKey string) error

	SessionAdd(ctx context.Context, sessionKey string, userID int) error
	SessionGet(ctx context.Context, sessionKey string) (userID int, err error)
	SessionDelete(ctx context.Context, sessionKey string) error

	ResetPasswordAdd(ctx context.Context, resetKey string, userID int) error
	ResetPasswordGet(ctx context.Context, resetKey string) (userID int, err error)
	ResetPasswordDelete(ctx context.Context, resetKey string) error
}
