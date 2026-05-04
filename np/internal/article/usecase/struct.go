package usecase

// =====================================================================
//  Query UseCase struct
// =====================================================================

type queryUseCase struct {
	queryRepository QueryRepository
	sessionProvider SessionProvider
}

// =====================================================================
//  Command UseCase struct
// =====================================================================

type commandUseCase struct {
	commandRepository CommandRepository
	threadInitialize  ThreadInitialize
	sessionProvider   SessionProvider
}
