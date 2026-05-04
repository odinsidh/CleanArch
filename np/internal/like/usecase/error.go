package usecase

import (
	"errors"
	"fmt"
)

// =====================================================================
//  custom UseCaseError
// =====================================================================

type UseCaseError struct {
	Input     any
	Operation string
	Err       error
}

func (u UseCaseError) Error() string {
	return fmt.Sprintf("%s : %v context(%v)", u.Operation, u.Err, u.Input)
}

func (u UseCaseError) Unwrap() error {
	return u.Err
}

// =====================================================================
//  Universal error
// =====================================================================

var (
	PermissionIssue error = errors.New("permission issue")
)

// =====================================================================
//  Query error
// =====================================================================

// =====================================================================
//  Command error
// =====================================================================
