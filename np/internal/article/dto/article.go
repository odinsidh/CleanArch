package dto

import "time"

type ArticleItem struct {
	ArticleID      int
	ArticleTitle   string
	ArticleContent string // short on slice view, full on per article request
	ArticleType    string
	AuthorID       int
	PublishedAt    time.Time
	ArticleActions ArticleActions
}

type ArticleActions struct {
	CanRead   bool
	CanModify bool
	CanDelete bool
}

func NewArticleItem(articleID int, articleTitle string, articleContent string, articleType string,
	authorID int, publishedAt time.Time, articleActions ArticleActions) *ArticleItem {
	return &ArticleItem{
		ArticleID:      articleID,
		ArticleTitle:   articleTitle,
		ArticleContent: articleContent,
		ArticleType:    articleType,
		AuthorID:       authorID,
		PublishedAt:    publishedAt,
		ArticleActions: articleActions,
	}
}
