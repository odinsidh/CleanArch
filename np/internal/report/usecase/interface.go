package usecase

import (
	"context"
	eAuthorization "newsportal/internal/authorization/entity"
	eReport "newsportal/internal/report/entity"
)

// =====================================================================
//  Another interfaces
// =====================================================================

type SessionProvider interface {
	Session(ctx context.Context, userID int) (*eAuthorization.Session, error)
}

// =====================================================================
//  user Query && Command Usecase | Repository
// =====================================================================

type UserQueryUsecase interface {
	GetCases(ctx context.Context, request Filter) ([]*eReport.Case, error)
	GetCaseByCaseID(ctx context.Context, userID, caseID int) (*eReport.Case, error)
}

type UserQueryRepository interface {
	GetCases(ctx context.Context, userID, cursor, offset int) ([]*eReport.Case, error)
	GetCaseByCaseID(ctx context.Context, userID, caseID int) (*eReport.Case, error)
}

type UserCommandUsecase interface {
	CreateReport(ctx context.Context, request CreateReport) error
}

type UserCommandRepository interface {
	CreateOrLoadCase(ctx context.Context, contentID int, contentSubID string, contentType, userID int) (*eReport.Case, error)
	CreateReport(ctx context.Context, enrichedCase *eReport.Case) error
}

// =====================================================================
//  moderator Query && Command Usecase | Repository
// =====================================================================

type ModeratorQueryUseCase interface {
	GetCases(ctx context.Context, request Filter) ([]*eReport.Case, error)
	GetMyCases(ctx context.Context, request Filter, status int) ([]*eReport.Case, error)
	GetCaseByID(ctx context.Context, moderatorID, caseID int) (*eReport.Case, error)
}

type ModeratorQueryRepository interface {
	GetCases(ctx context.Context, cursor, offset, sortType int) ([]*eReport.Case, error)
	GetCaseByID(ctx context.Context, moderatorID, caseID int) (*eReport.Case, error)
	GetCasesByModeratorID(ctx context.Context, moderatorID, cursor, offset, sortType, status int) ([]*eReport.Case, error)
}

type ModeratorCommandUseCase interface {
	TakeCaseByIDInWork(ctx context.Context, request CommandRequest) (*eReport.Case, error)
	AddCommentToCase(ctx context.Context, request CommandRequest) (*eReport.Case, error)
	MakeDecisionResolve(ctx context.Context, request CommandRequest) (*eReport.Case, error)
	MakeDecisionReject(ctx context.Context, request CommandRequest) (*eReport.Case, error)
	TransferCaseToAnotherModerator(ctx context.Context, request TransferCase) (*eReport.Case, error)
}

type ModeratorCommandRepository interface {
	Save(ctx context.Context, inputCase *eReport.Case) error
}
