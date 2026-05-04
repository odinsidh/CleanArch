package entity

import (
	"errors"
	"sort"
)

type Session struct {
	UserID           int
	Roles            []RoleContainer
	MergedPermission map[Domain]map[Permission]bool
}

var (
	ErrSessionUserIDIsEqualToZero error = errors.New("session, user id is equal to zero")
)

func LoadSession(userID int, roles []RoleContainer) (*Session, error) {
	if userID == 0 {
		return nil, ErrSessionUserIDIsEqualToZero
	}

	var activeRoles []RoleContainer = extractActiveRoles(roles)
	var permission map[Domain]map[Permission]bool = mergeAllRoles(activeRoles)

	output := Session{
		UserID:           userID,
		Roles:            roles,
		MergedPermission: permission,
	}

	return &output, nil
}

func extractActiveRoles(roles []RoleContainer) []RoleContainer {
	var container map[Role]RoleContainer = make(map[Role]RoleContainer)
	var output []RoleContainer

	sort.Slice(roles, func(i, j int) bool {
		return roles[i].Timestamp.Before(roles[j].Timestamp)
	})

	for index := range roles {
		currentRole := roles[index]
		if currentRole.EventType == Grant {
			container[currentRole.Role] = currentRole
			continue
		}
		if currentRole.EventType == Revoke {
			delete(container, currentRole.Role)
			continue
		}
	}

	for _, value := range container {
		output = append(output, value)
	}

	return output
}

func ExtractActiveRolesMap(roles []RoleContainer) map[Role]RoleContainer {
	var container map[Role]RoleContainer = make(map[Role]RoleContainer)

	sort.Slice(roles, func(i, j int) bool {
		return roles[i].Timestamp.Before(roles[j].Timestamp)
	})

	for index := range roles {
		currentRole := roles[index]
		if currentRole.EventType == Grant {
			container[currentRole.Role] = currentRole
			continue
		}
		if currentRole.EventType == Revoke {
			delete(container, currentRole.Role)
			continue
		}
	}

	return container
}

func mergeAllRoles(roles []RoleContainer) map[Domain]map[Permission]bool {
	var output map[Domain]map[Permission]bool = make(map[Domain]map[Permission]bool)

	for index := range roles {
		currentRole := roles[index].Role
		permissionPerRole := permissionControl[currentRole]

		for domain, permissionMap := range permissionPerRole {
			if output[domain] == nil {
				output[domain] = make(map[Permission]bool)
			}

			for permissionKey, permissionValue := range permissionMap {
				if permissionValue == true {
					output[domain][permissionKey] = true
				}
			}
		}
	}

	return output
}

func (self *Session) Can(domain Domain, permission Permission) bool {
	found, ok := self.MergedPermission[domain][permission]
	if ok {
		return found
	}
	return false
}
