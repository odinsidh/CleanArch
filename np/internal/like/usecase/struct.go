package usecase

// =====================================================================
//  universal service structures
// =====================================================================

type Deps struct {
	SessionProvider   SessionProvider
	QueryRepo         QueryRepository
	CommandRepository CommandRepository
}

// =====================================================================
//  Query && Command struct
// =====================================================================

type queryUseCase struct {
	sessionProvider SessionProvider
	queryRepo       QueryRepository
}

type commandUseCase struct {
	sessionProvider   SessionProvider
	commandRepository CommandRepository
}
