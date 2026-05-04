package entity

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrCommentWasAlreadyDeactivated error = fmt.Errorf("core : [domain, entity, Comment] error: [comment was already deactivated]")
	ErrModifyDeactivatedComment     error = errors.New("deleted comment cannot be modyfied")
	ErrUserIDMismatch               error = errors.New("missmatch request userID and owner userID")
	ErrPermissionIssue              error = errors.New("permission issue")
)

type Comment struct {
	TargetType     targetType
	TargetID       int
	MessageID      int
	ReplyMessageID int
	UserID         int
	CreatedAt      time.Time
	ModifiedAt     time.Time
	Value          value
	Active         bool
}

func NewComment(targetType string, targetID int, replyMessageID int, userID int, value string) (*Comment, error) {
	commentTargetType, err := NewTargetType(targetType)
	if err != nil {
		return nil, err
	}
	commentValue, err := NewValue(value)
	if err != nil {
		return nil, err
	}

	return &Comment{
		TargetType:     *commentTargetType,
		TargetID:       targetID,
		ReplyMessageID: replyMessageID,
		UserID:         userID,
		CreatedAt:      time.Now(),
		ModifiedAt:     time.Now(),
		Value:          *commentValue,
		Active:         true,
	}, nil
}

func (self *Comment) Deactivate(requestUserID int, globalPermission bool) error {
	if !self.CanBeDeletedByID(requestUserID, globalPermission) {
		return ErrPermissionIssue
	}
	err := self.deactivate()
	if err != nil {
		return err
	}

	return nil
}

func (self *Comment) deactivate() error {
	if self.Active == false {
		return ErrCommentWasAlreadyDeactivated
	}
	self.Active = false
	return nil
}

func (self *Comment) CanBeDeletedByID(userID int, globalPermission bool) bool {
	if self.UserID == userID || globalPermission == true {
		return true
	}

	return false
}

func (self *Comment) IsActive() bool {
	return self.Active
}

func (self *Comment) Update(requestUserID int, updatedComment string, globalPermission bool) error {
	if !self.IsActive() {
		return ErrModifyDeactivatedComment
	}

	self.CanBeUpdateByID(requestUserID, globalPermission)

	if self.UserID != requestUserID {
		return ErrUserIDMismatch
	}

	newComment, err := NewValue(updatedComment)
	if err != nil {
		return err
	}
	self.Value = *newComment

	return nil
}

func (self *Comment) CanBeUpdateByID(userID int, globalPermission bool) bool {
	if self.UserID == userID || globalPermission == true {
		return true
	}

	return false
}
