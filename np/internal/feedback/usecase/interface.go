package usecase

import (
	"context"
	eAuthorization "newsportal/internal/authorization/entity"
	eFeedback "newsportal/internal/feedback/entity"
)

// =====================================================================
//  Another interfaces
// =====================================================================

type SessionProvider interface {
	Session(ctx context.Context, userID int) (*eAuthorization.Session, error)
}

// =====================================================================
//  Query && Command Usecase | Repository
// =====================================================================

type QueryUseCase interface {
	LoadFeedbackWall(ctx context.Context, userID int, requestType string, cursor, amount int) ([]eFeedback.Feedback, error)
	LoadFeedbackByID(ctx context.Context, userID int, requestType string, feedbackID int) (*eFeedback.Feedback, error)
}

type QueryRepository interface {
	LoadFeedbackWall(ctx context.Context, requestType string, cursor, amount int) ([]eFeedback.Feedback, error)
	LoadFeedbackByID(ctx context.Context, requestType string, feedbackID int) (*eFeedback.Feedback, error)
}

type CommandUseCase interface {
	NewFeedback(ctx context.Context, userID int, email, requestType, message string) error
}

type CommandRepository interface {
	NewFeedback(ctx context.Context, feedback eFeedback.Feedback) error
	// CloseFeedback (просто метод закрытия по ID, + userID кто закрываел, + причина закрытия) автоматом выставляется дата закрытия
	// UpdateFeedback (тут типо комментарий добавляет человек по запросу, соответственно userID кто обновил)
	// ChangeStatus (userID who changed, updateTime)
}
