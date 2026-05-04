package usecase

import (
	"context"
	articleDTO "newsportal/internal/article/dto"
	eAuthorization "newsportal/internal/authorization/entity"
	feedDTO "newsportal/internal/feed/dto"
)

// =====================================================================
//  Another interfaces
// =====================================================================

type SessionProvider interface {
	Session(ctx context.Context, userID int) (*eAuthorization.Session, error)
}

// =====================================================================
//  Base
// =====================================================================

type ArticleQuery interface {
	GetArticles(ctx context.Context, filter Filter) ([]articleDTO.ArticleItem, error)
}

type CommentsQuery interface {
	GetCommentsCountByID(ctx context.Context, articleType string, articleID ...int) (map[int]int, error)
}

type UserQuery interface {
	GetUsernamesByID(ctx context.Context, userID ...int) (map[int]string, error)
}

// =====================================================================
//  UseCase && Repository
// =====================================================================

type Feed interface {
	GetFeed(ctx context.Context, filter Filter) (*feedDTO.FeedDTO, error)
}
