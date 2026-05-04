package usecase

import (
	"context"
	"fmt"
	eAuthorization "newsportal/internal/authorization/entity"
	eReport "newsportal/internal/report/entity"
)

func (self *moderatorCommandUseCase) TakeCaseByIDInWork(ctx context.Context, request CommandRequest) (*eReport.Case, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "report > moderatorCommand > TakeCaseByIDInWork"
	input := fmt.Sprintf("request: [%v]", request)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if request.ModeratorID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}

	if request.CaseID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrCaseIDIsEqualToZero, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Report
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead, eAuthorization.CanModify}
	_, err := havePermission(ctx, self.sessionProvider, request.ModeratorID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	outputCase, err := self.queryRepo.GetCaseByID(ctx, request.ModeratorID, request.CaseID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	err = outputCase.TakeInWork(request.ModeratorID, request.Message)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	err = self.commandRepo.Save(ctx, outputCase)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return outputCase, nil
}

func (self *moderatorCommandUseCase) AddCommentToCase(ctx context.Context, request CommandRequest) (*eReport.Case, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "report > moderatorCommand > AddCommentToCase"
	input := fmt.Sprintf("request: [%v]", request)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if request.ModeratorID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}

	if request.CaseID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrCaseIDIsEqualToZero, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Report
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead, eAuthorization.CanModify}
	_, err := havePermission(ctx, self.sessionProvider, request.ModeratorID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	outputCase, err := self.queryRepo.GetCaseByID(ctx, request.ModeratorID, request.CaseID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input,
		)
	}

	err = outputCase.AddModeratorMessage(request.ModeratorID, request.Message)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input,
		)
	}

	err = self.commandRepo.Save(ctx, outputCase)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input,
		)
	}

	return outputCase, nil
}

func (self *moderatorCommandUseCase) MakeDecisionResolve(ctx context.Context, request CommandRequest) (*eReport.Case, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "report > moderatorCommand > MakeDecisionResolve"
	input := fmt.Sprintf("request: [%v]", request)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if request.ModeratorID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}

	if request.CaseID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrCaseIDIsEqualToZero, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Report
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead, eAuthorization.CanModify}
	_, err := havePermission(ctx, self.sessionProvider, request.ModeratorID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	outputCase, err := self.queryRepo.GetCaseByID(ctx, request.ModeratorID, request.CaseID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	err = outputCase.AddModeratorDecisionResolve(request.ModeratorID, request.Message)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	err = self.commandRepo.Save(ctx, outputCase)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return outputCase, nil
}

func (self *moderatorCommandUseCase) MakeDecisionReject(ctx context.Context, request CommandRequest) (*eReport.Case, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "report > moderatorCommand > MakeDecisionReject"
	input := fmt.Sprintf("request: [%v]", request)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if request.ModeratorID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}

	if request.CaseID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrCaseIDIsEqualToZero, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Report
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead, eAuthorization.CanModify}
	_, err := havePermission(ctx, self.sessionProvider, request.ModeratorID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	outputCase, err := self.queryRepo.GetCaseByID(ctx, request.ModeratorID, request.CaseID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	err = outputCase.AddModeratorDecisionReject(request.ModeratorID, request.Message)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	err = self.commandRepo.Save(ctx, outputCase)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return outputCase, nil
}

func (self *moderatorCommandUseCase) TransferCaseToAnotherModerator(ctx context.Context, request TransferCase) (*eReport.Case, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "report > moderatorCommand > TransferCaseToAnotherModerator"
	input := fmt.Sprintf("request: [%v]", request)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if request.OldModeratorID == 0 || request.NewModeratorID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}

	if request.CaseID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrCaseIDIsEqualToZero, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Report
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead, eAuthorization.CanModify}
	_, err := havePermission(ctx, self.sessionProvider, request.OldModeratorID, permissionDomain, permissionToCheck...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	outputCase, err := self.queryRepo.GetCaseByID(ctx, request.OldModeratorID, request.CaseID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	err = outputCase.TransferCaseToAnotherModerator(request.OldModeratorID, request.NewModeratorID, request.Message)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	err = self.commandRepo.Save(ctx, outputCase)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return outputCase, nil
}
