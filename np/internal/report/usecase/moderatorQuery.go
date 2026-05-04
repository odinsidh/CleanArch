package usecase

import (
	"context"
	"fmt"
	eAuthorization "newsportal/internal/authorization/entity"
	eReport "newsportal/internal/report/entity"
)

func (self *moderatorQueryUseCase) GetCases(ctx context.Context, request Filter) ([]*eReport.Case, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "report > moderatorQuery > GetCases"
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
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead, eAuthorization.CanModify}
	_, err := havePermission(ctx, self.sessionProvider, request.UserID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	_, err = eReport.NewSortType(request.SortType)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	outputCases, err := self.queryRepo.GetCases(ctx, request.Cursor, request.Offset, request.SortType)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return outputCases, nil
}

func (self *moderatorQueryUseCase) GetMyCases(ctx context.Context, request Filter, status int) ([]*eReport.Case, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "report > moderatorQuery > GetMyCases"
	input := fmt.Sprintf("request: [%v] status: [%v]", request, status)

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
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead, eAuthorization.CanModify}
	_, err := havePermission(ctx, self.sessionProvider, request.UserID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	_, err = eReport.NewSortType(request.SortType)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	_, err = eReport.NewStatus(status)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	outputCases, err := self.queryRepo.GetCasesByModeratorID(ctx, request.UserID, request.Cursor, request.Offset, request.SortType, status)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return outputCases, nil
}

func (self *moderatorQueryUseCase) GetCaseByID(ctx context.Context, moderatorID, caseID int) (*eReport.Case, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "report > moderatorQuery > GetCaseByID"
	input := fmt.Sprintf("moderatorID: [%v] caseID: [%v]", moderatorID, caseID)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if moderatorID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Report
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead, eAuthorization.CanModify}
	_, err := havePermission(ctx, self.sessionProvider, moderatorID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	outputCase, err := self.queryRepo.GetCaseByID(ctx, moderatorID, caseID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return outputCase, nil
}
