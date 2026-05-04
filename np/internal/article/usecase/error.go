package usecase

import "errors"

// =====================================================================
//  Universal
// =====================================================================

var (
	ErrUserIDIsEqualToZero    error = errors.New("user id is equal to zero")
	ErrUserDontHavePermission error = errors.New("user dont have permission")
)

// =====================================================================
//  Query UseCase
// =====================================================================

var (
	ErrArticleTypeNotFound   error = errors.New("article type not found")
	ErrSortTypeIsEqualToZero error = errors.New("sort type is equal to zero")
)

// =====================================================================
//  Query Repository
// =====================================================================
