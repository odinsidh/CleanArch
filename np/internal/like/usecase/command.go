package usecase

import (
	"context"
	"fmt"
	eAuth "newsportal/internal/authorization/entity"
	dLike "newsportal/internal/like/dto"
	eLike "newsportal/internal/like/entity"
)

func (c *commandUseCase) hitCaller(ctx context.Context, request dLike.Reaction, reactionType eLike.ReactionType) error {
	// =====================================================================
	//  error tracing
	// =====================================================================
	op := fmt.Sprintf("like > commandUseCase > hitCaller > Hit%v", reactionType)

	// =====================================================================
	//  session
	// =====================================================================
	session, err := c.sessionProvider.Session(ctx, request.UserID)
	if err != nil {
		return errWrap(request, op, err)
	}

	domain := eAuth.Like
	permission := eAuth.CanRead
	if !session.Can(domain, permission) {
		return errWrap(request, op, PermissionError{
			Domain:     domain,
			Permission: permission,
		})
	}

	// =====================================================================
	//  input validation
	// =====================================================================
	reaction, err := eLike.NewReaction(reactionType, request.UserID, request.TargetID, request.TargetType)
	if err != nil {
		return errWrap(request, op, err)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	// check if reaction already exist
	repoReaction, err := c.commandRepository.IsReactionExist(ctx, reaction)
	if err != nil {
		return errWrap(request, op, err)
	}

	// if exist call touch reaction
	if repoReaction != nil {
		reaction = repoReaction

		// call touch only if reaction already exist
		switch reactionType {
		case eLike.Positive:
			reaction.TouchPositive()
		case eLike.Negative:
			reaction.TouchNegative()
		}
	}

	// save new state to repository
	err = c.commandRepository.ApplyReaction(ctx, *reaction)
	if err != nil {
		return errWrap(request, op, err)
	}

	return nil
}

func (c *commandUseCase) HitPositive(ctx context.Context, request dLike.Reaction) error {
	return c.hitCaller(ctx, request, eLike.Positive)
}

func (c *commandUseCase) HitNegative(ctx context.Context, request dLike.Reaction) error {
	return c.hitCaller(ctx, request, eLike.Negative)
}
