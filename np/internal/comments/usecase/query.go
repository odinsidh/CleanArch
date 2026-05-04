package usecase

import (
	"context"
	"fmt"
	eAuthorization "newsportal/internal/authorization/entity"
	eComments "newsportal/internal/comments/entity"
)

func (self *queryUseCase) GetCommentsCountByID(ctx context.Context, userID int, targetType string, targetID ...int) (map[int]int, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "comments > queryUseCase > GetCommentsCountByID"
	input := fmt.Sprintf("userID: [%v] targetType: [%v] targetID: [%v]", userID, targetType, targetID)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if userID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Messages
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead}
	_, err := havePermission(ctx, self.sessionProvider, userID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	if targetType == "" {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrTargetTypeNotFound, input)
	}

	// for the future OR LEN IS MORE THAN 100
	if len(targetID) == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrTargetIDLenIsEqualToZero, input)
	}
	// for the future, cache something on REDIS (reduce courier work on dat level)

	commentCount, err := self.queryRepo.GetCommentsCountByID(ctx, targetType, targetID...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return commentCount, nil
}

func (self *queryUseCase) GetComments(ctx context.Context, userID int, selector Selector, cursor, pagination int) ([]*eComments.Comment, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "comments > queryUseCase > GetComments"
	input := fmt.Sprintf("userID: [%v] selector: [%v] cursor: [%v] pagination: [%v]", userID, selector, cursor, pagination)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if userID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Messages
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead}
	_, err := havePermission(ctx, self.sessionProvider, userID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	comments, err := self.queryRepo.GetComments(ctx, selector, cursor, pagination)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return comments, nil
}
