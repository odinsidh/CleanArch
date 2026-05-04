package usecase

import (
	"context"
	"fmt"
	eAuthorization "newsportal/internal/authorization/entity"
	eFeedback "newsportal/internal/feedback/entity"
)

func (self *commandUseCase) NewFeedback(ctx context.Context, userID int, email, requestType, message string) error {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "feedback > commandUseCase > NewFeedback"
	input := fmt.Sprintf("userID: [%v] email: [%v] requestType: [%v] message: [%v]", userID, email, requestType, message)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if userID == 0 {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Feedback
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead, eAuthorization.CanCreate}
	_, err := havePermission(ctx, self.sessionProvider, userID, permissionDomain, permissionToCheck...)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	feedback, err := eFeedback.NewFeedback(userID, email, requestType, message)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	err = self.commandRepo.NewFeedback(ctx, *feedback)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return nil
}
