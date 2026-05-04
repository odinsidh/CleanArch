package dto

import "errors"

var (
	ErrAttachUsernameNotFound error = errors.New("feed dto issue, pair not found, on authodID and usermane")
	ErrAttachCommentNotFound  error = errors.New("feed dto issue, pair not found, on articleID and comment count")
)
