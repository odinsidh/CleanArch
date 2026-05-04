package dto

import (
	"time"
)

type FeedDTO struct {
	FeedItems []FeedItemDTO
}

type FeedItemDTO struct {
	ArticleID      int
	ArticleTitle   string
	ArticleContent string // short on slice view, full on per article request
	ArticleType    string
	AuthorID       int
	Timestamp      time.Time
	ArticleActions ArticleActions
	AuthorUsername string
	CommentsCount  int
}

type ArticleActions struct {
	CanRead   bool
	CanModify bool
	CanDelete bool
}

func (self *FeedDTO) AttachUsername(linkedAuthorIDAndUsernames map[int]string) error {
	for index := range self.FeedItems {
		authorID := self.FeedItems[index].AuthorID

		username, ok := linkedAuthorIDAndUsernames[authorID]
		if !ok {
			return ErrAttachUsernameNotFound
		}
		self.FeedItems[index].AuthorUsername = username
	}

	return nil
}

func (self *FeedDTO) AttachCommentAmount(linkedArticleIDandCommentAmount map[int]int) error {
	for index := range self.FeedItems {
		articleID := self.FeedItems[index].ArticleID

		commentCount, ok := linkedArticleIDandCommentAmount[articleID]
		if !ok {
			return ErrAttachCommentNotFound
		}
		self.FeedItems[index].CommentsCount = commentCount
	}

	return nil
}
