package usecase

import (
	"context"
	"fmt"
	eAuthorization "newsportal/internal/authorization/entity"
)

func (self *userCommandUseCase) CreateReport(ctx context.Context, request CreateReport) error {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "report > userCommand > CreateReport"
	input := fmt.Sprintf("request: [%v] ", request)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if request.UserID == 0 {
		return fmt.Errorf("%s : %w context(%v)",
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
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	// Создаем или загружаем кейс
	reportCase, err := self.commandRepo.CreateOrLoadCase(ctx, request.ContentID, request.ContentSubID, request.ContentType, request.UserID)
	if err != nil {
		return fmt.Errorf("%s :%w context(%v)",
			operation, err, input,
		)
	}

	// пытаемся присоединить жалобу к Case
	err = reportCase.AddReport(request.UserID, request.UserMessage, request.UserReason)
	if err != nil {
		return fmt.Errorf("%s :%w context(%v)",
			operation, err, input,
		)
	}

	// отдаем CASE базе
	err = self.commandRepo.CreateReport(ctx, reportCase)
	if err != nil {
		return fmt.Errorf("%s :%w context(%v)",
			operation, err, input,
		)
	}

	return nil
}
