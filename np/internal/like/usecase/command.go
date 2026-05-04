package usecase

import (
	"context"
	"fmt"
	eAuth "newsportal/internal/authorization/entity"
	dLike "newsportal/internal/like/dto"
	eLike "newsportal/internal/like/entity"
)

func (c *commandUseCase) HitPositive(ctx context.Context, request dLike.Reaction) error {
	// =====================================================================
	//  error tracing
	// =====================================================================
	op := "like > commandUseCase > HitPositive"

	// =====================================================================
	//  session
	// =====================================================================
	session, err := c.sessionProvider.Session(ctx, request.UserID)
	if err != nil {
		return UseCaseError{
			Input:     request,
			Operation: op,
			Err:       err,
		}
	}

	domain := eAuth.Like
	permission := eAuth.CanRead
	if !session.Can(domain, permission) {
		return UseCaseError{
			Input:     request,
			Operation: op,
			Err:       fmt.Errorf("err(%w) domain(%v) permission(%v)", PermissionIssue, domain, permission),
		}

	}

	// =====================================================================
	//  input validation
	// =====================================================================
	reactionType := eLike.Positive
	reaction, err := eLike.NewReaction(int(reactionType), request.UserID, request.TargetID, request.TargetType)
	if err != nil {
		return UseCaseError{
			Input:     request,
			Operation: op,
			Err:       err,
		}
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	// check if reaction already exist
	repoReaction, err := c.commandRepository.IsReactionExist(ctx, request.UserID, request.TargetID, request.TargetType)
	if err != nil {
		return UseCaseError{
			Input:     request,
			Operation: op,
			Err:       err,
		}
	}

	// if exist - reaatach
	if repoReaction != nil {
		reaction = repoReaction
		// call touch only if reaction already exist
		reaction.TouchPositive()
	}

	// save new state to repository
	err = c.commandRepository.ApplyReaction(ctx, *reaction)
	if err != nil {
		return UseCaseError{
			Input:     request,
			Operation: op,
			Err:       err,
		}
	}

	return nil
}

func (c *commandUseCase) HitNegative(ctx context.Context, request dLike.Reaction) error {
	// =====================================================================
	//  error tracing
	// =====================================================================
	op := "like > commandUseCase > HitNegative"

	// =====================================================================
	//  session
	// =====================================================================
	session, err := c.sessionProvider.Session(ctx, request.UserID)
	if err != nil {
		return UseCaseError{
			Input:     request,
			Operation: op,
			Err:       err,
		}
	}

	domain := eAuth.Like
	permission := eAuth.CanRead
	if !session.Can(domain, permission) {
		return UseCaseError{
			Input:     request,
			Operation: op,
			Err:       fmt.Errorf("err(%w) domain(%v) permission(%v)", PermissionIssue, domain, permission),
		}

	}

	// =====================================================================
	//  input validation
	// =====================================================================
	reactionType := eLike.Negative
	reaction, err := eLike.NewReaction(int(reactionType), request.UserID, request.TargetID, request.TargetType)
	if err != nil {
		return UseCaseError{
			Input:     request,
			Operation: op,
			Err:       err,
		}
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	// check if reaction already exist
	repoReaction, err := c.commandRepository.IsReactionExist(ctx, request.UserID, request.TargetID, request.TargetType)
	if err != nil {
		return UseCaseError{
			Input:     request,
			Operation: op,
			Err:       err,
		}
	}

	// if exist - reaatach
	if repoReaction != nil {
		reaction = repoReaction
		// call touch only if reaction already exist
		reaction.TouchNegative()
	}

	// save new state to repository
	err = c.commandRepository.ApplyReaction(ctx, *reaction)
	if err != nil {
		return UseCaseError{
			Input:     request,
			Operation: op,
			Err:       err,
		}
	}

	return nil
}
