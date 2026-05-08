package usecase

import (
	"context"
	dLike "newsportal/internal/like/dto"
)

// =====================================================================
//	Query interface
// =====================================================================

type QueryUseCase interface {
	CountReaction(ctx context.Context, request dLike.CountRequest) ([]CountResponse, error)
}

// =====================================================================
//  Command interface
// =====================================================================

type CommandUseCase interface {
	HitPositive(ctx context.Context, request dLike.Reaction) error
	HitNegative(ctx context.Context, request dLike.Reaction) error
}
