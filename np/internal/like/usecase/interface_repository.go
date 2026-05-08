package usecase

import (
	"context"
	eAuthorization "newsportal/internal/authorization/entity"
	eLike "newsportal/internal/like/entity"
)

//go:generate mockgen -source=interface_repository.go -destination=../mock/mock_test.go -package=mockgen

// =====================================================================
//  Universal interface
// =====================================================================

type SessionProvider interface {
	Session(ctx context.Context, userID int) (*eAuthorization.Session, error)
}

// =====================================================================
//	Query interface
// =====================================================================

type QueryRepository interface {
	CountReaction(ctx context.Context, request []CountRequest) ([]CountResponse, error)
}

// =====================================================================
//  Command interface
// =====================================================================

type CommandRepository interface {
	IsReactionExist(ctx context.Context, request *eLike.Reaction) (*eLike.Reaction, error)
	ApplyReaction(ctx context.Context, reaction eLike.Reaction) error
}
