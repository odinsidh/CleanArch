package usecase

// =====================================================================
//  another struct
// =====================================================================

type Deps struct {
	commandRepo      CommandRepository
	commandRepoCache CommandRepositoryCache
	queryRepo        QueryRepository
	queryRepoCache   QueryRepositoryCache
	notification     Notification
}

type UserIDAndPassword struct {
	UserID   int
	Password string
}

// =====================================================================
//  Query struct
// =====================================================================

type queryUseCase struct {
	queryRepo      QueryRepository
	queryRepoCache QueryRepositoryCache
}

// =====================================================================
//  Command struct
// =====================================================================

type commandUseCase struct {
	commandRepo      CommandRepository
	commandRepoCache CommandRepositoryCache
	queryRepo        QueryRepository
	queryRepoCache   QueryRepositoryCache
	notification     Notification
}
