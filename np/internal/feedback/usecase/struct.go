package usecase

// =====================================================================
//  universal service structures
// =====================================================================

// =====================================================================
//  Query && Command struct
// =====================================================================

type queryUseCase struct {
	queryRepo       QueryRepository
	sessionProvider SessionProvider
}

type commandUseCase struct {
	commandRepo     CommandRepository
	sessionProvider SessionProvider
}
