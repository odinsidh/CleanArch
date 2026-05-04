package usecase

import (
	"context"
	"fmt"
	eAuthorization "newsportal/internal/authorization/entity"
	eUser "newsportal/internal/user/entity"
)

func (self *queryUseCase) GetUser(ctx context.Context, requestFromUserID, targetUserID int) (*eUser.User, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "user > queryUseCase > GetUser"
	input := fmt.Sprintf("requestUserID: [%v] targetUserID: [%v]", requestFromUserID, targetUserID)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if requestFromUserID == 0 || targetUserID == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDsLenIsEqualToZero, input)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	session, err := self.sessionProvider.Session(ctx, requestFromUserID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	permissionDomain := eAuthorization.Users
	permissionRead := eAuthorization.CanRead
	if !session.Can(permissionDomain, permissionRead) {
		return nil, fmt.Errorf("%s :%w context(%v) domain: [%v] permission: [%v]",
			operation, ErrUserDontHavePermission, input, permissionDomain, permissionRead)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	user, err := self.queryRepository.GetUser(ctx, targetUserID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return user, nil
}

func (self *queryUseCase) GetUsernamesByID(ctx context.Context, requestFromUserID int, userIDs ...int) (map[int]string, error) {
	// TODO LIMITATION MAX LENGHT userIDs

	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "user > queryUseCase > GetUsernamesByID"
	input := fmt.Sprintf("userID: [%v] userIDs: [%v]", requestFromUserID, userIDs)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if requestFromUserID == 0 || len(userIDs) == 0 {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDsLenIsEqualToZero, input)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	session, err := self.sessionProvider.Session(ctx, requestFromUserID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	permissionDomain := eAuthorization.Users
	permissionRead := eAuthorization.CanRead
	if !session.Can(permissionDomain, permissionRead) {
		return nil, fmt.Errorf("%s : %w context(%v) domain: [%v] permission: [%v]",
			operation, ErrUserDontHavePermission, input, permissionDomain, permissionRead)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	// redis cache will be live here
	usernames, err := self.queryRepository.GetUsernamesByID(ctx, userIDs...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return usernames, nil
}
