package usecase

import "time"

// =====================================================================
//  Query struct
// =====================================================================

type queryUseCase struct {
	sessionProvider SessionProvider
	queryRepository QueryRepository
}

// =====================================================================
//  Command struct
// =====================================================================

type commandUseCase struct {
	sessionProvider    SessionProvider
	commandRepository  CommandRepository
	ownQueryRepository OwnQueryRepository
}

type CreateUser struct {
	Username string
	Email    string
	Password string
	Age      time.Time
}
