package dto

type CreateUser struct {
	Username string
	Email    string
	Password string
}

type Login struct {
	Username string
	Password string
}

type Logout struct {
	SessionKey string
}

type ResetRequire struct {
	Username string
}

type ResetConfirm struct {
	ResetID     string
	NewPassword string
}

type EmailConfirm struct {
	RegisterKey string
}

type Cookie struct {
	SessionKey string
}

type ChangePassword struct {
	SessionKey  string
	OldPassword string
	NewPassword string
}
type ChangeUsername struct {
	SessionKey  string
	NewUsername string
}

type ChangeEmail struct {
	SessionKey string
	NewEmail   string
}
