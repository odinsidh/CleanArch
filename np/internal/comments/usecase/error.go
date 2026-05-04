package usecase

import "errors"

// =====================================================================
//  Universal
// =====================================================================

var (
	ErrUserIDIsEqualToZero    error = errors.New("user id is equal to zero")
	ErrUserDontHavePermission error = errors.New("user dont have a permission")
)

// =====================================================================
//  user Query && CommandUseCase
// =====================================================================

var (
	ErrTargetTypeNotFound       error = errors.New("article type not found")
	ErrTargetIDIsEqualToZero    error = errors.New("target id is equal to zero")
	ErrTargetIDLenIsEqualToZero error = errors.New("target id lenght is equal to zero")
	ErrMessageIDIsEqualToZero   error = errors.New("messageID is equal to zero")
)

var (
	ErrReplyCommentNotFound                                error = errors.New("reply comment was not found")
	ErrTryingToAddCommentWithNotEqualThreadIDandThreadType error = errors.New("attemp to add comment to different threadID and threadType")
)
