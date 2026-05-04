package usecase

import (
	"context"
	"fmt"
	eAuthN "newsportal/internal/authentication/entity"
)

func (self *queryUseCase) GetUserIDByUsername(ctx context.Context, username string) (userID int, err error) {
	// =====================================================================
	//  error tracing
	// =====================================================================
	var operation string = "authentication > queryUseCase > GetUserIDByUsername"
	var input string = fmt.Sprintf("username: [%v]", username)

	// =====================================================================
	//  usecase checkers
	// =====================================================================

	// =====================================================================
	//  core logick
	// =====================================================================
	_, err = eAuthN.NewUsername(username)
	if err != nil {
		return userID, &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	userID, err = self.queryRepo.UserID(ctx, username)
	if err != nil {
		return userID, &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	return userID, nil
}

func (self *queryUseCase) DoesUsernameAvaliable(ctx context.Context, username string) error {
	// =====================================================================
	//  pre define variables
	// =====================================================================
	var (
		err     error
		isTaken bool
	)

	// =====================================================================
	//  error tracing
	// =====================================================================
	var operation string = "authentication > queryUseCase > DoesUsernameAvaliable"
	var input string = fmt.Sprintf("username: [%v]", username)

	// =====================================================================
	//  usecase checkers
	// =====================================================================

	// =====================================================================
	//  core logick
	// =====================================================================
	_, err = eAuthN.NewUsername(username)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	isTaken, err = self.queryRepoCache.UsernameIsExist(ctx, username)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}
	if isTaken {
		return ErrUsernameAlreadyExist
	}

	isTaken, err = self.queryRepo.UsernameIsExist(ctx, username)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}
	if isTaken {
		return ErrUsernameAlreadyExist
	}

	return nil
}
