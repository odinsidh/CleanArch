package usecase

import (
	"context"
	eAuth "newsportal/internal/authorization/entity"
	dLike "newsportal/internal/like/dto"
	eLike "newsportal/internal/like/entity"
)

func (c *queryUseCase) CountReaction(ctx context.Context, request dLike.CountRequest) ([]CountResponse, error) {
	// =====================================================================
	// 	error tracing
	// =====================================================================
	op := "like > queryUseCase > CountReaction"

	// =====================================================================
	// 	checkers
	// =====================================================================
	if request.UserID <= 0 {
		return nil, &UseCaseError{
			Input:     request,
			Operation: op,
			Err:       UserIDEqualToZero,
		}
	}
	if len(request.Targets) == 0 {
		return nil, &UseCaseError{
			Input:     request,
			Operation: op,
			Err:       TargetLengthIsEqualToZero,
		}
	}

	// =====================================================================
	//	session
	// =====================================================================
	session, err := c.sessionProvider.Session(ctx, request.UserID)
	if err != nil {
		return nil, errWrap(request, op, err)
	}

	domain := eAuth.Like
	permission := eAuth.CanRead
	if !session.Can(domain, permission) {
		return nil, errWrap(request, op, &PermissionError{
			Domain:     domain,
			Permission: permission,
		})
	}

	// =====================================================================
	//	input validation
	// =====================================================================
	targets := make([]CountRequest, 0, len(request.Targets))

	for index := range request.Targets {

		validTargetID, err := eLike.NewTargetID(request.Targets[index].TargetID)
		if err != nil {
			return nil, errWrap(request, op, &ValidationError{
				Input: request.Targets[index].TargetID,
				Err:   err,
			})
		}

		validTargetType, err := eLike.NewTargetType(request.Targets[index].TargetType)
		if err != nil {
			return nil, errWrap(request, op, &ValidationError{
				Input: request.Targets[index].TargetType,
				Err:   err,
			})
		}

		targets = append(targets, CountRequest{
			Target: Target{
				ID:   validTargetID,
				Type: validTargetType,
			},
		})

	}

	// =====================================================================
	//	core logick
	// =====================================================================

	output, err := c.queryRepo.CountReaction(ctx, targets)
	if err != nil {
		return nil, errWrap(request, op, err)
	}

	return output, nil
}
