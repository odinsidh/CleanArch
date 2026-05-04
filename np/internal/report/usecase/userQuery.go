package usecase

import (
	"context"
	"fmt"
	eAuthorization "newsportal/internal/authorization/entity"
	eReport "newsportal/internal/report/entity"
)

func (self *userQueryUseCase) GetCases(ctx context.Context, request Filter) ([]*eReport.Case, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "report > userQuery > GetCases"
	input := fmt.Sprintf("request: [%v]", request)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if request.UserID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Report
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead}
	_, err := havePermission(ctx, self.sessionProvider, request.UserID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	outputCases, err := self.queryRepo.GetCases(ctx, request.UserID, request.Cursor, request.Offset)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return outputCases, nil
}

func (self *userQueryUseCase) GetCaseByCaseID(ctx context.Context, userID, caseID int) (*eReport.Case, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "report > userQuery > GetCaseByCaseID"
	input := fmt.Sprintf("userID: [%v] caseID: [%v]", userID, caseID)

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
	permissionDomain := eAuthorization.Report
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead}
	_, err := havePermission(ctx, self.sessionProvider, userID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	outputCase, err := self.queryRepo.GetCaseByCaseID(ctx, userID, caseID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return outputCase, nil
}
