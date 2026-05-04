package usecase

func NewQueryUseCase(queryRepository QueryRepository, sessionProvider SessionProvider) QueryUseCase {
	return &queryUseCase{
		queryRepository: queryRepository,
		sessionProvider: sessionProvider,
	}
}

func NewCommandUseCase(commandRepository CommandRepository, threadInitialize ThreadInitialize, sessionProvider SessionProvider) CommandUseCase {
	return &commandUseCase{
		commandRepository: commandRepository,
		threadInitialize:  threadInitialize,
		sessionProvider:   sessionProvider,
	}
}
