package usecase

// =====================================================================
//  Query && Command constructor
// =====================================================================

func NewQueryUseCase(queryRepo QueryRepository, sessionProvider SessionProvider) QueryUseCase {

	return &queryUseCase{
		queryRepo:       queryRepo,
		sessionProvider: sessionProvider,
	}
}

func NewCommandUseCase(commandRepo CommandRepository, sessionProvider SessionProvider) CommandUseCase {

	return &commandUseCase{
		commandRepo:     commandRepo,
		sessionProvider: sessionProvider,
	}
}
