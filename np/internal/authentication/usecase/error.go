package usecase

import "fmt"

// =====================================================================
//  custom UseCaseError
// =====================================================================

type UseCaseError struct {
	Operation string
	Input     any
	Err       error
}

func (self *UseCaseError) Error() string {
	return fmt.Sprintf("%s : %v context(%v)", self.Operation, self.Err, self.Input)
}

func (self *UseCaseError) Unwrap() error {
	return self.Err
}

// =====================================================================
//  Universal
// =====================================================================

// =====================================================================
//  Query UseCase
// =====================================================================

var (
	ErrUsernameAlreadyExist error = fmt.Errorf("username already exist")
	ErrUUIDEmpty            error = fmt.Errorf("generated uuid was empty")
	ErrCoreFieldInDTOEmpty  error = fmt.Errorf("core dto field come empty")
	ErrInvalidPassword      error = fmt.Errorf("invalid password")
)

// =====================================================================
//  Command UseCase
// =====================================================================
