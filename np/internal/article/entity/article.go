package entity

import (
	"errors"
	"time"
)

var (
	ErrArticleUpdateRightIssue  error = errors.New("user have no rights to update this article")
	ErrArticleDeleteRightIssue  error = errors.New("user have no rights to delete this article")
	ErrArticlePublishRightIssue error = errors.New("user have no rights to publish this article")
)

type Article struct {
	AuthorID    int
	ArticleID   int
	ArticleType string
	Title       title
	Content     content
	Status      status
	CreatedAT   time.Time
	UpdatedAT   time.Time
	Active      bool
}

func NewArticle(authorID int, articleType string, title string, content string) (*Article, error) {
	newTitle, err := NewTitle(title)
	if err != nil {
		return nil, err
	}
	newContent, err := NewContent(content)
	if err != nil {
		return nil, err
	}
	newStatus, err := NewStatus()
	if err != nil {
		return nil, err
	}

	output := Article{
		AuthorID:    authorID,
		ArticleType: articleType,
		Title:       *newTitle,
		Content:     *newContent,
		Status:      *newStatus,
		CreatedAT:   time.Now(),
		UpdatedAT:   time.Now(),
		Active:      true,
	}
	return &output, nil
}

func (self *Article) Publish() {
	self.Status.Publish()
}

func (self *Article) DeleteArticle() {
	self.Active = false
}

func (self *Article) CanBeUpdatedByID(userID int, sessionPermission bool) bool {
	if sessionPermission || self.AuthorID == userID {
		return true
	}

	return false
}

func (self *Article) CanBeDeletedByID(userID int, sessionPermission bool) bool {
	if sessionPermission || self.AuthorID == userID {
		return true
	}

	return false
}

func (self *Article) CanBePublishByID(userID int, sessionPermission bool) bool {
	if sessionPermission || self.AuthorID == userID {
		return true
	}

	return false
}

func (self *Article) UpdateArticle(newTitle string, newContent string) error {
	var err error
	err = self.updateTitle(newTitle)
	if err != nil {
		return err
	}

	err = self.updateContent(newContent)
	if err != nil {
		return err
	}

	self.updateUpdatedAT()
	return nil
}

func (self *Article) updateTitle(newTitle string) error {
	updatedTitle, err := NewTitle(newTitle)
	if err != nil {
		return err
	}
	self.Title = *updatedTitle
	return nil
}

func (self *Article) updateContent(newContent string) error {
	updatedContent, err := NewContent(newContent)
	if err != nil {
		return err
	}
	self.Content = *updatedContent
	return nil
}

func (self *Article) updateUpdatedAT() {
	self.UpdatedAT = time.Now()
}
