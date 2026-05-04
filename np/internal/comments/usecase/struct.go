package usecase

// =====================================================================
//  universal service structures
// =====================================================================

type Selector struct {
	TargetType string
	TargetID   int
}

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
