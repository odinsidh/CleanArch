package usecase

// =====================================================================
//  Query constructor
// =====================================================================

func NewQueryUseCase(sessionProvider SessionProvider, queryRepository QueryRepository) QueryUseCase {
	return &queryUseCase{
		queryRepository: queryRepository,
		sessionProvider: sessionProvider,
	}
}

// =====================================================================
//  Command constructor
// =====================================================================

func NewCommandUseCase(sessionProvider SessionProvider, commandRepository CommandRepository, ownQueryRepository OwnQueryRepository) CommandUseCase {
	return &commandUseCase{
		sessionProvider:    sessionProvider,
		commandRepository:  commandRepository,
		ownQueryRepository: ownQueryRepository,
	}
}
