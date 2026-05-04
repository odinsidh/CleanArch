package entity

import (
	"errors"
	"fmt"
	"time"
)

type RoleContainer struct {
	RoleContainerID int
	Role            Role
	EventType       EventType
	ToUserID        int
	ByUserID        int
	Message         string
	Timestamp       time.Time
}

var (
	ErrRoleContainerUserIDIsEqualToZero error = errors.New("RoleContainer, userID is equal to zero")
	ErrRoleContainerMessageNotFound     error = errors.New("RoleContainer, assigned message not found")
	ErrRoleContainerMessageTooLong      error = errors.New("RoleContainer, assigned message too long")
	ErrRolesLenghtIsEqualToZero         error = errors.New("RoleContainer, roles lenght is equal to zero")
)

func RevokeRole(role Role, toUserID, byUserID int, message string) (*RoleContainer, error) {
	eventType := Revoke
	return newRoleContainer(role, eventType, toUserID, byUserID, message)
}

func GrantRole(role Role, toUserID, byUserID int, message string) (*RoleContainer, error) {
	eventType := Grant
	return newRoleContainer(role, eventType, toUserID, byUserID, message)
}

func GrantRoles(roles []Role, toUserID, byUserID int, message string) ([]RoleContainer, error) {
	eventType := Grant
	if len(roles) == 0 {
		return nil, fmt.Errorf("%w | event type: [%v] ", ErrRolesLenghtIsEqualToZero, eventType)
	}

	var output []RoleContainer = make([]RoleContainer, 0, len(roles))

	for index := range roles {
		currentRole := roles[index]
		roleContainer, err := newRoleContainer(currentRole, eventType, toUserID, byUserID, message)
		if err != nil {
			return nil, err
		}
		output = append(output, *roleContainer)
	}

	return output, nil
}

func RevokeRoles(roles []Role, toUserID, byUserID int, message string) ([]RoleContainer, error) {
	eventType := Revoke
	if len(roles) == 0 {
		return nil, fmt.Errorf("%w | event type: [%v] ", ErrRolesLenghtIsEqualToZero, eventType)
	}

	var output []RoleContainer = make([]RoleContainer, 0, len(roles))

	for index := range roles {
		currentRole := roles[index]
		roleContainer, err := newRoleContainer(currentRole, eventType, toUserID, byUserID, message)
		if err != nil {
			return nil, err
		}
		output = append(output, *roleContainer)
	}

	return output, nil
}

func newRoleContainer(role Role, eventType EventType, toUserID, byUserID int, message string) (*RoleContainer, error) {
	if toUserID == 0 || byUserID == 0 {
		return nil, ErrRoleContainerUserIDIsEqualToZero
	}
	if message == "" {
		return nil, ErrRoleContainerMessageNotFound
	}
	if len(message) > 500 {
		return nil, ErrRoleContainerMessageTooLong
	}
	output := RoleContainer{
		Role:      role,
		EventType: eventType,
		ToUserID:  toUserID,
		ByUserID:  byUserID,
		Message:   message,
		Timestamp: time.Now(),
	}

	return &output, nil
}
