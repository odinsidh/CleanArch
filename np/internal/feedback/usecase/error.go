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
