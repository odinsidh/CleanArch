package usecase

// =====================================================================
//  user Query && Command constructor
// =====================================================================

func NewUserQueryUsecase(queryRepo UserQueryRepository,
	sessionProvider SessionProvider) UserQueryUsecase {

	return &userQueryUseCase{
		queryRepo:       queryRepo,
		sessionProvider: sessionProvider,
	}
}

func NewUserCommandUsecase(commandRepo UserCommandRepository,
	sessionProvider SessionProvider) UserCommandUsecase {

	return &userCommandUseCase{
		commandRepo:     commandRepo,
		sessionProvider: sessionProvider,
	}
}

// =====================================================================
//  moderator Query && Command constructor
// =====================================================================

func NewModeratorQueryUseCase(queryRepo ModeratorQueryRepository,
	sessionProvider SessionProvider) ModeratorQueryUseCase {

	return &moderatorQueryUseCase{
		queryRepo:       queryRepo,
		sessionProvider: sessionProvider,
	}
}

func NewModeratorCommandUseCase(commandRepo ModeratorCommandRepository,
	queryRepo ModeratorQueryRepository,
	sessionProvider SessionProvider) ModeratorCommandUseCase {

	return &moderatorCommandUseCase{
		commandRepo:     commandRepo,
		queryRepo:       queryRepo,
		sessionProvider: sessionProvider,
	}
}
