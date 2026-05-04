package usecase

// =====================================================================
//  Query constructor
// =====================================================================

func NewQueryUseCase(deps Deps) QueryUseCase {
	return &queryUseCase{
		queryRepo:      deps.queryRepo,
		queryRepoCache: deps.queryRepoCache,
	}
}

// =====================================================================
//  Command constructor
// =====================================================================

func NewCommandUseCase(deps Deps) CommandUseCase {
	return &commandUseCase{
		commandRepo:      deps.commandRepo,
		commandRepoCache: deps.commandRepoCache,
		queryRepo:        deps.queryRepo,
		queryRepoCache:   deps.queryRepoCache,
		notification:     deps.notification,
	}
}
