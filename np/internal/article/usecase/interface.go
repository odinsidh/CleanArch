package usecase

import (
	"context"
	"newsportal/internal/article/dto"
	eArticle "newsportal/internal/article/entity"
	eAuthorization "newsportal/internal/authorization/entity"
)

// =====================================================================
//  Another interfaces
// =====================================================================

type SessionProvider interface {
	Session(ctx context.Context, userID int) (*eAuthorization.Session, error)
}

type ThreadInitialize interface {
	Initialize(ctx context.Context, threadType string, threadID int) error
}

// =====================================================================
//  Query UseCase && Query Repository
// =====================================================================

type QueryUseCase interface {
	GetArticles(ctx context.Context, userId int, articleType string, cursor int, limit int, sortType int) ([]dto.ArticleItem, error)
	GetArticle(ctx context.Context, userID int, articleType string, articleID int) (*dto.ArticleItem, error)
}

type QueryRepository interface {
	GetArticles(ctx context.Context, articleType string, cursor int, limit int, sortType int) ([]dto.ArticleItem, error)
	GetArticle(ctx context.Context, targetType string, targetID int) (*eArticle.Article, error)
}

// =====================================================================
//  Command UseCase && Command Repository
// =====================================================================

type CommandUseCase interface {
	CreateArticle(ctx context.Context, authorID int, articleType string, title string, content string) (*dto.ArticleItem, error)
	ModifyArticle(ctx context.Context, authorID int, articleType string, articleID int, title string, content string) (*dto.ArticleItem, error)
	DeleteArticle(ctx context.Context, authorID int, articleType string, articleID int) error
	PublishArticle(ctx context.Context, authorID int, articleType string, articleID int) error
}

type CommandRepository interface {
	CreateArticle(ctx context.Context, input *eArticle.Article) (*eArticle.Article, error)
	ReadArticle(ctx context.Context, targetType string, targetID int) (*eArticle.Article, error)
	ModifyArticle(ctx context.Context, input *eArticle.Article) (*eArticle.Article, error)
	DeleteArticle(ctx context.Context, input *eArticle.Article) error
	PublishArticle(ctx context.Context, input *eArticle.Article) error
	HardDelete(ctx context.Context, input *eArticle.Article) error
}
