package usecase

import "errors"

// =====================================================================
//  Universal
// =====================================================================

var (
	ErrUserDontHavePermission error = errors.New("user dont have permission")
)

// =====================================================================
//  Query UseCase
// =====================================================================

var (
	ErrUserIDsLenIsEqualToZero error = errors.New("user IDs lenght is equal to zero")
)

// =====================================================================
//  Command UseCase
// =====================================================================

var (
	ErrUserAlreadyExist error = errors.New("user already exist")
)
