package usecase

import (
	"context"
	"fmt"
	eAuthorization "newsportal/internal/authorization/entity"
	eFeedback "newsportal/internal/feedback/entity"
)

func (self *queryUseCase) LoadFeedbackWall(ctx context.Context, userID int, requestType string, cursor, amount int) ([]eFeedback.Feedback, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "feedback > queryUseCase > LoadFeedbackWall"
	input := fmt.Sprintf("userID: [%v] requestType: [%v] cursor: [%v] amount: [%v]", userID, requestType, cursor, amount)

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
	permissionDomain := eAuthorization.Feedback
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead}
	_, err := havePermission(ctx, self.sessionProvider, userID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	feedbacks, err := self.queryRepo.LoadFeedbackWall(ctx, requestType, cursor, amount)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// тут может жить redis, но это фидбеки, не то чтобы оно должно было грузить базу сильно
	return feedbacks, nil
}

func (self *queryUseCase) LoadFeedbackByID(ctx context.Context, userID int, requestType string, feedbackID int) (*eFeedback.Feedback, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "feedback > queryUseCase > LoadFeedbackByID"
	input := fmt.Sprintf("userID: [%v] requestType: [%v] feedbackID: [%v]", userID, requestType, feedbackID)

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
	permissionDomain := eAuthorization.Feedback
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead}
	_, err := havePermission(ctx, self.sessionProvider, userID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	feedback, err := self.queryRepo.LoadFeedbackByID(ctx, requestType, feedbackID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// тут может жить redis, но это фидбеки, не то чтобы оно должно было грузить базу сильно
	return feedback, nil
}
