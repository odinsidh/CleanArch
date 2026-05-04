package entity

type Role int

const (
	Moderator Role = iota + 1
	Administrator
	Ownership
	User
)

func AvailableRoles() []Role {
	return []Role{Administrator, Moderator, Ownership}
}
