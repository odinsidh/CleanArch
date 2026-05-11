package usecase

import (
	"errors"
	"fmt"

	eAuthorization "newsportal/internal/authorization/entity"
)

// =====================================================================
//  custom UseCaseError
// =====================================================================

type UseCaseError struct {
	Input     any
	Operation string
	Err       error
}

func (u *UseCaseError) Error() string {
	return fmt.Sprintf("%s : %v context(%v)", u.Operation, u.Err, u.Input)
}

func (u *UseCaseError) Unwrap() error {
	return u.Err
}

func errWrap(input any, operation string, err error) *UseCaseError {
	return &UseCaseError{
		Input:     input,
		Operation: operation,
		Err:       err,
	}
}

// =====================================================================
//  Permission error
// =====================================================================

var (
	PermissionIssue error = errors.New("permission issue")
)

type PermissionError struct {
	Domain     eAuthorization.Domain
	Permission eAuthorization.Permission
}

func (p *PermissionError) Error() string {
	return fmt.Sprintf("permission issue, domain(%v) permission(%v)", p.Domain, p.Permission)
}

func (p *PermissionError) Unwrap() error {
	return PermissionIssue
}

// =====================================================================
//	Validation error
// =====================================================================

var (
	ValidateIssue error = errors.New("validate issue")
)

type ValidationError struct {
	Input any
	Err   error
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%v, input(%v) validate error(%v)", ValidateIssue, v.Input, v.Err)
}

func (v *ValidationError) Unwrap() error {
	return v.Err
}

// =====================================================================
//  Universal error
// =====================================================================

var (
	UserIDEqualToZero error = errors.New("userID is equal to zero")
)

// =====================================================================
//	Query error
// =====================================================================

var (
	TargetLengthIsEqualToZero error = errors.New("target length is equal to zero")
)

// =====================================================================
//  Command error
// =====================================================================
