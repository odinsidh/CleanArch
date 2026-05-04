package usecase

// =====================================================================
//  Query && Command cobstructor
// =====================================================================

func NewQueryUseCase(deps *Deps) QueryUseCase {
	return &queryUseCase{
		sessionProvider: deps.SessionProvider,
		queryRepo:       deps.QueryRepo,
	}
}

func NewCommandUseCase(deps *Deps) CommandUseCase {
	return &commandUseCase{
		sessionProvider:   deps.SessionProvider,
		commandRepository: deps.CommandRepository,
	}
}
