package entity

import "errors"

// =====================================================================
//  errors
// =====================================================================

var (
	ErrTargetIDNotValid       error = errors.New("target id not valid")
	ErrTargetTypeNotFound     error = errors.New("target type not found")
	ErrTargetTypeNotAvailable error = errors.New("target type not available")
	ErrUserIDNotFound         error = errors.New("user id not found")
	ErrReactionTypeNotExist   error = errors.New("reaction type not exist")
)
