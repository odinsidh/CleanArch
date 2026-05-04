package usecase

import (
	"context"
	"errors"
	"fmt"
	"newsportal/internal/authorization/entity"
)

var (
	ErrRolesLenEqualToZero                            error = errors.New("roles len equal to zero")
	ErrUserDontHaveActiveRoleToRevoke                 error = errors.New("user dont have active role to revoke")
	ErrRevokeRolesInputRoleAmountNotEqualToFactAmount error = errors.New("input amount of roles not equal to fact amount of roles")
)

type CommandRepo interface {
	GrantRole(ctx context.Context, roleContainer entity.RoleContainer) error
	GrantRoles(ctx context.Context, roleContainer []entity.RoleContainer) error
	RevokeRole(ctx context.Context, roleContainer entity.RoleContainer) error
	RevokeRoles(ctx context.Context, roleContainer []entity.RoleContainer) error
	GetRoles(ctx context.Context, userID int) ([]entity.RoleContainer, error)
}

type CommandUseCase struct {
	commandRepo     CommandRepo
	sessionProvider SessionProvider
}

type SessionProvider interface {
	Session(ctx context.Context, userID int) (*entity.Session, error)
}

func NewCommandUseCase(commandRepo CommandRepo, sessionProvider SessionProvider) *CommandUseCase {
	return &CommandUseCase{
		commandRepo:     commandRepo,
		sessionProvider: sessionProvider,
	}
}

// Назначение ролей
func (self *CommandUseCase) AddRole(ctx context.Context, role entity.Role, byUserID, toUserID int, message string) error {
	// input validation
	if byUserID == 0 || toUserID == 0 {
		return fmt.Errorf("%w | CommandUseCase > AddRole | byUserID: [%v] toUserID: [%v]", ErrUserIDIsEqualToZero, byUserID, toUserID)
	}

	// auth validation
	session, err := self.sessionProvider.Session(ctx, byUserID)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > AddRole | session create issue for UserID: [%v]", err, byUserID)
	}

	domain := entity.Authorization
	permissionModify := entity.CanModify
	if !session.Can(domain, permissionModify) {
		return fmt.Errorf("%w | CommandUseCase > AddRole > Session.Can | failed combination > domain: [%v] permission: [%v]", ErrPermissionIssue, domain, permissionModify)
	}

	// create roleContainer
	roleContainer, err := entity.GrantRole(role, toUserID, byUserID, message)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > AddRole > Entity.GrantRole | failed combination > role: [%v] byUserID: [%v] toUserID: [%v] message: [%v]", err, role, byUserID, toUserID, message)
	}

	// save roleContainer
	err = self.commandRepo.GrantRole(ctx, *roleContainer)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > AddRole > commandRepo.AddRole", err)
	}

	return nil
}

func (self *CommandUseCase) AddRoles(ctx context.Context, roles []entity.Role, byUserID, toUserID int, message string) error {
	// input validation
	if byUserID == 0 || toUserID == 0 {
		return fmt.Errorf("%w | CommandUseCase > AddRoles | byUserID: [%v] toUserID: [%v]", ErrUserIDIsEqualToZero, byUserID, toUserID)
	}

	if len(roles) == 0 {
		return fmt.Errorf("%w | CommandUseCase > AddRoles", ErrRolesLenEqualToZero)
	}

	// auth validation
	session, err := self.sessionProvider.Session(ctx, byUserID)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > AddRoles | session create issue for UserID: [%v]", err, byUserID)
	}

	domain := entity.Authorization
	permissionModify := entity.CanModify
	if !session.Can(domain, permissionModify) {
		return fmt.Errorf("%w | CommandUseCase > AddRoles > Session.Can | failed combination > domain: [%v] permission: [%v]", ErrPermissionIssue, domain, permissionModify)
	}

	// create roleContainer
	rolesContainer, err := entity.GrantRoles(roles, toUserID, byUserID, message)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > AddRoles > Entity.GrantRoles | failed > roles: [%v] byUserID: [%v] toUserID: [%v] message: [%v]", err, roles, byUserID, toUserID, message)
	}

	// save roleContainer
	err = self.commandRepo.GrantRoles(ctx, rolesContainer)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > AddRoles > commandRepo.AddRoles", err)
	}

	return nil
}

func (self *CommandUseCase) AddAllRoles(ctx context.Context, byUserID, toUserID int, message string) error {
	// input validation
	if byUserID == 0 || toUserID == 0 {
		return fmt.Errorf("%w | CommandUseCase > AddAllRoles | byUserID: [%v] toUserID: [%v]", ErrUserIDIsEqualToZero, byUserID, toUserID)
	}

	// auth validation
	session, err := self.sessionProvider.Session(ctx, byUserID)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > AddAllRoles | session create issue for UserID: [%v]", err, byUserID)
	}

	domain := entity.Authorization
	permissionRead := entity.CanRead
	permissionModify := entity.CanModify

	if !session.Can(domain, permissionRead) {
		return fmt.Errorf("%w | CommandUseCase > AddAllRoles > Session.Can | failed combination > domain: [%v] permission: [%v]", ErrPermissionIssue, domain, permissionRead)
	}
	if !session.Can(domain, permissionModify) {
		return fmt.Errorf("%w | CommandUseCase > AddAllRoles > Session.Can | failed combination > domain: [%v] permission: [%v]", ErrPermissionIssue, domain, permissionModify)
	}

	// take all avaliable roles
	roles := entity.AvailableRoles()

	// create roleContainer
	rolesContainer, err := entity.GrantRoles(roles, toUserID, byUserID, message)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > AddAllRoles > Entity.GrantRoles | failed > roles: [%v] byUserID: [%v] toUserID: [%v] message: [%v]", err, roles, byUserID, toUserID, message)
	}

	// save roleContainer
	err = self.commandRepo.GrantRoles(ctx, rolesContainer)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > AddAllRoles > commandRepo.AddRoles", err)
	}

	return nil
}

// Снятие ролей
func (self *CommandUseCase) RevokeRole(ctx context.Context, role entity.Role, byUserID, toUserID int, message string) error {
	// input validation
	if byUserID == 0 || toUserID == 0 {
		return fmt.Errorf("%w | CommandUseCase > RevokeRole | byUserID: [%v] toUserID: [%v]", ErrUserIDIsEqualToZero, byUserID, toUserID)
	}

	// auth validation
	session, err := self.sessionProvider.Session(ctx, byUserID)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > RevokeRole | session create issue for UserID: [%v]", err, byUserID)
	}

	domain := entity.Authorization
	permissionRead := entity.CanRead
	permissionModify := entity.CanModify

	if !session.Can(domain, permissionRead) {
		return fmt.Errorf("%w | CommandUseCase > RevokeRole > Session.Can | failed combination > domain: [%v] permission: [%v]", ErrPermissionIssue, domain, permissionRead)
	}
	if !session.Can(domain, permissionModify) {
		return fmt.Errorf("%w | CommandUseCase > RevokeRole > Session.Can | failed combination > domain: [%v] permission: [%v]", ErrPermissionIssue, domain, permissionModify)
	}

	// читаем существующие роли
	rolesContainer, err := self.commandRepo.GetRoles(ctx, toUserID)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > RevokeRole > commandRepo.GetRoles | failed > toUserId: [%v]", err, toUserID)
	}

	// извлекаем все активные
	activeRoles := entity.ExtractActiveRolesMap(rolesContainer)

	// проверяем есть ли среди активной нужная по запросу, отклоняем в случае ошибки
	_, ok := activeRoles[role]
	if !ok {
		return ErrUserDontHaveActiveRoleToRevoke
	}

	// исполняем (создаем обьект)
	revokeRoleContainer, err := entity.RevokeRole(role, toUserID, byUserID, message)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > RevokeRole > entity.RevokeRole | failed > toUserId: [%v] byUserID: [%v] message: [%v]",
			err, toUserID, byUserID, message)
	}

	// пишем в базу
	err = self.commandRepo.RevokeRole(ctx, *revokeRoleContainer)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > RevokeRole > commandRepo.RevokeRole |", err)
	}

	return nil
}

func (self *CommandUseCase) RevokeRoles(ctx context.Context, roles []entity.Role, byUserID, toUserID int, message string) error {
	// input validation
	if byUserID == 0 || toUserID == 0 {
		return fmt.Errorf("%w | CommandUseCase > RevokeRoles | byUserID: [%v] toUserID: [%v]", ErrUserIDIsEqualToZero, byUserID, toUserID)
	}

	if len(roles) == 0 {
		return fmt.Errorf("%w | CommandUseCase > RevokeRoles", ErrRolesLenEqualToZero)
	}

	// auth validation
	session, err := self.sessionProvider.Session(ctx, byUserID)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > RevokeRoles | session create issue for UserID: [%v]", err, byUserID)
	}

	domain := entity.Authorization
	permissionRead := entity.CanRead
	permissionModify := entity.CanModify

	if !session.Can(domain, permissionRead) {
		return fmt.Errorf("%w | CommandUseCase > RevokeRoles > Session.Can | failed combination > domain: [%v] permission: [%v]", ErrPermissionIssue, domain, permissionRead)
	}
	if !session.Can(domain, permissionModify) {
		return fmt.Errorf("%w | CommandUseCase > RevokeRoles > Session.Can | failed combination > domain: [%v] permission: [%v]", ErrPermissionIssue, domain, permissionModify)
	}

	// читаем существующие роли
	rolesContainer, err := self.commandRepo.GetRoles(ctx, toUserID)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > RevokeRoles > commandRepo.GetRoles | failed > toUserId: [%v]", err, toUserID)
	}

	// revokeRoles Checker
	// извлекаем все активные
	activeRoles := entity.ExtractActiveRolesMap(rolesContainer)

	// защищаем ввод от дублей
	var mergedRoles map[entity.Role]struct{} = make(map[entity.Role]struct{})
	for index := range roles {
		curentRole := roles[index]
		mergedRoles[curentRole] = struct{}{}
	}

	for roleToRevoke := range mergedRoles {
		_, ok := activeRoles[roleToRevoke]
		if !ok {
			return fmt.Errorf("%w | CommandUseCase > RevokeRoles > revokeRoles checker | failed > toUserID: [%v]", ErrUserDontHaveActiveRoleToRevoke, toUserID)
		}
	}

	// собираем роли для отзыва
	var rolesToRevoke []entity.Role
	for key := range mergedRoles {
		rolesToRevoke = append(rolesToRevoke, key)
	}

	// создаем entity
	revokeRoleContainer, err := entity.RevokeRoles(rolesToRevoke, toUserID, byUserID, message)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > RevokeRoles > entity.RevokeRole | failed > toUserId: [%v] byUserID: [%v] message: [%v]",
			err, toUserID, byUserID, message)
	}

	// пишем в репозиторий
	err = self.commandRepo.RevokeRoles(ctx, revokeRoleContainer)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > RevokeRoles > commandRepo.RevokeRole |", err)
	}

	return nil
}

func (self *CommandUseCase) RevokeAllRoles(ctx context.Context, byUserID, toUserID int, message string) error {
	// input validation
	if byUserID == 0 || toUserID == 0 {
		return fmt.Errorf("%w | CommandUseCase > RevokeAllRoles | byUserID: [%v] toUserID: [%v]", ErrUserIDIsEqualToZero, byUserID, toUserID)
	}

	// auth validation
	session, err := self.sessionProvider.Session(ctx, byUserID)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > RevokeAllRoles | session create issue for UserID: [%v]", err, byUserID)
	}

	domain := entity.Authorization
	permissionRead := entity.CanRead
	permissionModify := entity.CanModify

	if !session.Can(domain, permissionRead) {
		return fmt.Errorf("%w | CommandUseCase > RevokeAllRoles > Session.Can | failed combination > domain: [%v] permission: [%v]", ErrPermissionIssue, domain, permissionRead)
	}
	if !session.Can(domain, permissionModify) {
		return fmt.Errorf("%w | CommandUseCase > RevokeAllRoles > Session.Can | failed combination > domain: [%v] permission: [%v]", ErrPermissionIssue, domain, permissionModify)
	}

	// Читаем существующие роли у пользователя
	roleContainer, err := self.commandRepo.GetRoles(ctx, toUserID)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > RevokeAllRoles > commandRepo.GetRoles | failed > userID: [%v]", err, toUserID)
	}

	// выносим только активные
	activeRoles := entity.ExtractActiveRolesMap(roleContainer)
	if len(activeRoles) == 0 {
		return ErrUserDontHaveActiveRoleToRevoke
	}

	// map to slice
	var activeRolesSlice []entity.Role = make([]entity.Role, 0, len(activeRoles))
	for key := range activeRoles {
		activeRolesSlice = append(activeRolesSlice, key)
	}

	// создаем entity
	revokeRoleContainer, err := entity.RevokeRoles(activeRolesSlice, toUserID, byUserID, message)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > RevokeAllRoles > entity.RevokeRole | failed > toUserId: [%v] byUserID: [%v] message: [%v]",
			err, toUserID, byUserID, message)
	}

	// пишем в репозиторий
	err = self.commandRepo.RevokeRoles(ctx, revokeRoleContainer)
	if err != nil {
		return fmt.Errorf("%w | CommandUseCase > RevokeAllRoles > commandRepo.RevokeRole |", err)
	}

	return nil
}

// Отредактировать permissionControl (да, для этого его нужно вынести в тот же json,
// и сохранять загружать), мы про это знаем помним, но пока скипаем
// но само собой мы это пока делать не будем, у нас учебные цели
func (self *CommandUseCase) ModifyRolePermission(ctx context.Context, byUserID int, newPermissions entity.PermissionControlSignature) error {
	// input validation
	var op string = " | CommandUseCase > EditRolePermission | "
	if byUserID == 0 {
		return fmt.Errorf("%w %v byUserID: [%v]", ErrUserIDIsEqualToZero, op, byUserID)
	}

	// auth validation
	session, err := self.sessionProvider.Session(ctx, byUserID)
	if err != nil {
		return fmt.Errorf("%w %v session create issue for UserID: [%v]", err, op, byUserID)
	}

	domain := entity.Authorization
	permissionRead := entity.CanRead
	permissionModify := entity.CanModify

	if !session.Can(domain, permissionRead) {
		return fmt.Errorf("%w %v failed combination > domain: [%v] permission: [%v]", ErrPermissionIssue, op, domain, permissionRead)
	}
	if !session.Can(domain, permissionModify) {
		return fmt.Errorf("%w %v failed combination > domain: [%v] permission: [%v]", ErrPermissionIssue, op, domain, permissionModify)
	}

	err = entity.ModifyRolePermission(newPermissions)
	if err != nil {
		return fmt.Errorf("%w %v failed > newPermissions: [%v]", err, op, newPermissions)
	}

	return nil
}
