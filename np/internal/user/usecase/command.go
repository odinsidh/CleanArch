package usecase

import (
	"context"
	"fmt"
	eAuthorization "newsportal/internal/authorization/entity"
	eUser "newsportal/internal/user/entity"
)

func (self *commandUseCase) CreateUser(ctx context.Context, request CreateUser) (*eUser.User, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "user > commandUseCase > CreateUser"
	input := fmt.Sprintf("request: [%v]", request)

	// =====================================================================
	//  core logick
	// =====================================================================
	user, err := eUser.NewUser(request.Username, request.Email, request.Password, request.Age)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// Проверить, не существует ли уже с такой почтой пользователь
	_, err = self.ownQueryRepository.GetUserByEmail(ctx, user.GetEmail())
	if err == nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserAlreadyExist, input)
	}

	// Здесь по-хорошему надо проверить, что err - это именно "Не найдено",
	// а не отвал базы данных. Если это отвал БД - возвращаем ошибку.
	// if !errors.Is(err, repository.ErrUserNotFound) {
	//     return nil, fmt.Errorf("%s : database error: %w", operation, err)
	// }

	// Создать пользователя
	createdUser, err := self.commandRepository.CreateUser(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return createdUser, nil

}

func (self *commandUseCase) DeactivateUser(ctx context.Context, requestFromUserID, targetUserID int) error {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "user > commandUseCase > DeactivateUser"
	input := fmt.Sprintf("requestFromUserID: [%v] targetUserID: [%v]", requestFromUserID, targetUserID)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if requestFromUserID == 0 || targetUserID == 0 {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDsLenIsEqualToZero, input)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	session, err := self.sessionProvider.Session(ctx, requestFromUserID)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	permissionDomain := eAuthorization.Users
	permissionRead := eAuthorization.CanRead
	if !session.Can(permissionDomain, permissionRead) {
		return fmt.Errorf("%s :%w context(%v) domain: [%v] permission: [%v]",
			operation, ErrUserDontHavePermission, input, permissionDomain, permissionRead)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	user, err := self.ownQueryRepository.GetUserById(ctx, targetUserID)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	permissionModify := eAuthorization.CanModify
	if !user.CanBeDeactivateByUser(requestFromUserID, session.Can(permissionDomain, permissionModify)) {
		return fmt.Errorf("%s :%w context(%v) domain: [%v] permission: [%v]",
			operation, ErrUserDontHavePermission, input, permissionDomain, permissionModify)
	}

	err = user.DeactivateUser()
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	err = self.commandRepository.DeactivateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return nil
}
