package usecase

import (
	"context"
	"errors"
	"fmt"
	"newsportal/internal/authorization/entity"
)

var (
	ErrUserIDIsEqualToZero error = errors.New("user id is equal to zero")
	ErrPermissionIssue     error = errors.New("permission issue")
)

type QueryRepo interface {
	Roles(ctx context.Context, userID int) ([]entity.RoleContainer, error)
}

type QueryUseCase struct {
	queryRepo QueryRepo
}

func NewQueryUseCase(queryRepo QueryRepo) *QueryUseCase {
	return &QueryUseCase{
		queryRepo: queryRepo,
	}
}

func (self *QueryUseCase) Session(ctx context.Context, userID int) (*entity.Session, error) {
	if userID == 0 {
		return nil, ErrUserIDIsEqualToZero
	}

	roles, err := self.queryRepo.Roles(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w, QueryUseCase > Session > queryRepo.Roles > input userID: [%v]", err, userID)
	}

	session, err := entity.LoadSession(userID, roles)
	if err != nil {
		return nil, fmt.Errorf("%w, QueryUseCase > Session > entity.LoadSession > input userID: [%v], roles: [%v]", err, userID, roles)
	}

	return session, nil
}

// Получить список существующих ролей
func (self *QueryUseCase) AvailableRoles(ctx context.Context, userID int) ([]entity.Role, error) {
	session, err := self.Session(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w, Authorization > QueryUseCase > Available Roles, input userID: [%v]", err, userID)
	}

	domain := entity.Authorization
	permission := entity.CanRead
	if !session.Can(domain, permission) {
		return nil, fmt.Errorf("%w, QueryUseCase > Available Roles, input values userID: [%v] domain: [%v] permission: [%v]", ErrPermissionIssue, userID, domain, permission)
	}

	availableRoles := entity.AvailableRoles()

	return availableRoles, nil
}

// Получить список существующих ролей
func (self *QueryUseCase) AvailablePermission(ctx context.Context, userID int) (entity.PermissionControlSignature, error) {
	session, err := self.Session(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w, Authorization > QueryUseCase > Available Roles, input userID: [%v]", err, userID)
	}

	domain := entity.Authorization
	permission := entity.CanRead
	if !session.Can(domain, permission) {
		return nil, fmt.Errorf("%w, QueryUseCase > Available Permission, input values userID: [%v] domain: [%v] permission: [%v]", ErrPermissionIssue, userID, domain, permission)
	}

	availablePermission := entity.AvailablePermission()

	return availablePermission, nil
}
