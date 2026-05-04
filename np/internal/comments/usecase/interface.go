package usecase

import (
	"context"
	eAuthorization "newsportal/internal/authorization/entity"
	eComments "newsportal/internal/comments/entity"
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
	GetCommentsCountByID(ctx context.Context, userID int, targetType string, targetID ...int) (map[int]int, error)
	GetComments(ctx context.Context, userID int, selector Selector, cursor, pagination int) ([]*eComments.Comment, error)
}

type QueryRepository interface {
	GetCommentsCountByID(ctx context.Context, targetType string, targetID ...int) (map[int]int, error)
	GetComments(ctx context.Context, selector Selector, cursor int, pagination int) ([]*eComments.Comment, error)
}

type CommandUseCase interface {
	AddComment(ctx context.Context, userID int, selector Selector, replyToID int, value string) error
	UpdateComment(ctx context.Context, userID int, selector Selector, messageID int, value string) error
	DeleteComment(ctx context.Context, userID int, selector Selector, messageID int) error
}

type CommandRepository interface {
	GetCommentByID(ctx context.Context, targetType string, targetID int, messageID int) (*eComments.Comment, error)
	CreateThread(ctx context.Context, selector Selector) error
	AddComment(ctx context.Context, comment *eComments.Comment) error
	DeleteComment(ctx context.Context, comment *eComments.Comment) error
	UpdateComment(ctx context.Context, comment *eComments.Comment) error
}
