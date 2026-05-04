package usecase

import (
	"context"
	eAuthorization "newsportal/internal/authorization/entity"
	dLike "newsportal/internal/like/dto"
	eLike "newsportal/internal/like/entity"
)

// =====================================================================
//  Universal interface
// =====================================================================

type SessionProvider interface {
	Session(ctx context.Context, userID int) (*eAuthorization.Session, error)
}

// =====================================================================
//	Query interface
// =====================================================================

type QueryUseCase interface {
	CountReaction(ctx context.Context, request []dLike.CountInput) ([]dLike.CountOutput, error)
}

type QueryRepository interface{}

// =====================================================================
//  Command interface
// =====================================================================

type CommandUseCase interface {
	// IsReactionExist(ctx context.Context, request dLike.Reaction) (*eLike.Reaction, error)
	HitPositive(ctx context.Context, request dLike.Reaction) error
	HitNegative(ctx context.Context, request dLike.Reaction) error
}

type CommandRepository interface {
	IsReactionExist(ctx context.Context, userID int, targetID int, targetType string) (*eLike.Reaction, error)
	ApplyReaction(ctx context.Context, reaction eLike.Reaction) error
}
