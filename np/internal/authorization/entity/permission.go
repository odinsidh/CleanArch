package entity

import (
	"errors"
	"sync"
)

var (
	ErrNewPermissionNotFound error = errors.New("new permissions is nil or equal to zero")
)

type Domain int

const (
	Article Domain = iota + 1
	Authorization
	Messages
	Feedback
	Report
	Users
	Feed
	Like
)

type Permission int

const (
	CanCreate Permission = iota + 1
	CanRead              // Предположим, что это отвечает за то,
	// что пользователь может в админке это видеть
	CanModify
	CanDelete
)

type PermissionControlSignature map[Role]map[Domain]map[Permission]bool

var permissionControlMutex sync.RWMutex = sync.RWMutex{}
var permissionControl PermissionControlSignature = PermissionControlSignature{

	User: map[Domain]map[Permission]bool{
		Article: map[Permission]bool{
			CanRead:   true,
			CanCreate: true,
		},
		Messages: map[Permission]bool{
			CanRead: true,
		},
		Feedback: map[Permission]bool{
			CanCreate: true,
		},
		Report: map[Permission]bool{
			CanCreate: true,
		},
		Users: map[Permission]bool{
			CanRead: true,
		},
		Feed: map[Permission]bool{
			CanRead: true,
		},
	},

	Moderator: map[Domain]map[Permission]bool{
		Messages: map[Permission]bool{
			CanRead:   true,
			CanDelete: true,
		},
		Feedback: map[Permission]bool{
			CanRead:   true,
			CanDelete: true,
		},
		Report: map[Permission]bool{
			CanRead:   true,
			CanDelete: true,
		},
		Users: map[Permission]bool{
			CanRead:   true,
			CanDelete: true,
		},
		Feed: map[Permission]bool{
			CanRead: true,
		},
	},

	Administrator: map[Domain]map[Permission]bool{
		Article: map[Permission]bool{
			CanCreate: true,
			CanRead:   true,
			CanModify: true,
			CanDelete: true,
		},
		Authorization: map[Permission]bool{
			CanCreate: true,
			CanRead:   true,
			CanModify: true,
			CanDelete: true,
		},
		Messages: map[Permission]bool{
			CanRead:   true,
			CanModify: true,
			CanDelete: true,
		},
		Feedback: map[Permission]bool{
			CanRead:   true,
			CanModify: true,
			CanDelete: true,
		},
		Report: map[Permission]bool{
			CanRead:   true,
			CanModify: true,
			CanDelete: true,
		},
		Users: map[Permission]bool{
			CanCreate: true,
			CanRead:   true,
			CanModify: true,
			CanDelete: true,
		},
		Feed: map[Permission]bool{
			CanRead: true,
		},
	},

	Ownership: map[Domain]map[Permission]bool{
		Article: map[Permission]bool{
			CanCreate: true,
			CanRead:   true,
			CanModify: true,
			CanDelete: true,
		},
		Messages: map[Permission]bool{
			CanCreate: true,
			CanRead:   true,
			CanModify: true,
			CanDelete: true,
		},
		Feedback: map[Permission]bool{
			CanCreate: true,
			CanRead:   true,
			CanModify: true,
			CanDelete: true,
		},
		Report: map[Permission]bool{
			CanCreate: true,
			CanRead:   true,
			CanModify: true,
			CanDelete: true,
		},
		Users: map[Permission]bool{
			CanCreate: true,
			CanRead:   true,
			CanModify: true,
			CanDelete: true,
		},
		Feed: map[Permission]bool{
			CanRead: true,
		},
	},
}

func AvailablePermission() PermissionControlSignature {
	var output PermissionControlSignature = make(PermissionControlSignature)

	permissionControlMutex.RLock()
	defer permissionControlMutex.RUnlock()

	for role, domainMap := range permissionControl {
		if output[role] == nil {
			output[role] = make(map[Domain]map[Permission]bool)
		}

		for domain, permissionMap := range domainMap {
			if output[role][domain] == nil {
				output[role][domain] = make(map[Permission]bool)
			}

			for permissionKey, permissionValue := range permissionMap {
				output[role][domain][permissionKey] = permissionValue
			}
		}
	}

	return output
}

func ModifyRolePermission(newPermissions PermissionControlSignature) error {
	if len(newPermissions) == 0 {
		return ErrNewPermissionNotFound
	}

	permissionControlMutex.Lock()
	defer permissionControlMutex.Unlock()

	for role, domainMap := range newPermissions {
		if permissionControl[role] == nil {
			permissionControl[role] = make(map[Domain]map[Permission]bool)
		}

		for domain, permissionMap := range domainMap {
			// убираем старые данные
			permissionControl[role][domain] = nil
			// новая инициализация
			if permissionControl[role][domain] == nil {
				permissionControl[role][domain] = make(map[Permission]bool)
			}

			for permissionKey, permissionValue := range permissionMap {
				permissionControl[role][domain][permissionKey] = permissionValue
			}
		}
	}

	return nil
}
