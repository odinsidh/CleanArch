package entity

import (
	"newsportal/internal/article/dto"
	"time"
)

type FeedItem struct {
	ArticleID             int
	ArticleTitle          string
	ArticlePreview        string
	ArticleAuthorUsername string
	PublishedAt           time.Time
	CommentsCount         int
}

func LoadFeedItem(articleID int, articleTitle string, articlePreview string, publishedAt time.Time) *FeedItem {
	output := FeedItem{
		ArticleID:      articleID,
		ArticleTitle:   articleTitle,
		ArticlePreview: articlePreview,
		PublishedAt:    publishedAt,
	}
	return &output
}

type Feed struct {
	FeedItems []FeedItem
}

func NewFeed(input []dto.ArticleItem) *Feed {
	output := Feed{}
	for _, value := range input {
		output.FeedItems = append(output.FeedItems, *LoadFeedItem(value.ArticleID, value.ArticleTitle, value.ArticleContent, value.PublishedAt))
	}

	return &output
}
