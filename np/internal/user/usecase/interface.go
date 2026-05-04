package usecase

import (
	"context"
	eAuthorization "newsportal/internal/authorization/entity"
	eUser "newsportal/internal/user/entity"
)

// =====================================================================
//  Another interfaces
// =====================================================================

type SessionProvider interface {
	Session(ctx context.Context, userID int) (*eAuthorization.Session, error)
}

// =====================================================================
//  Query UseCase && Repository
// =====================================================================

type QueryUseCase interface {
	GetUser(ctx context.Context, requestFromUserID, targetUserID int) (*eUser.User, error)
	GetUsernamesByID(ctx context.Context, requestFromUserID int, userIDs ...int) (map[int]string, error)
}

type QueryRepository interface {
	GetUser(ctx context.Context, userId int) (*eUser.User, error)
	GetUsernamesByID(ctx context.Context, userIDs ...int) (map[int]string, error)
}

// =====================================================================
//  Command UseCase && Repository
// =====================================================================

type CommandUseCase interface {
	CreateUser(ctx context.Context, createUser CreateUser) (*eUser.User, error)
	DeactivateUser(ctx context.Context, requestFromUserID, targetUserID int) error
}

type CommandRepository interface {
	CreateUser(ctx context.Context, createUser CreateUser) (*eUser.User, error) // здесь я возвращаю dto.User,
	DeactivateUser(ctx context.Context, user *eUser.User) error
}

type OwnQueryRepository interface {
	GetUserById(ctx context.Context, userId int) (*eUser.User, error)          // helper for deactivate
	GetUserByEmail(ctx context.Context, userEmail string) (*eUser.User, error) // helper for creation
}
