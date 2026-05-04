package usecase

// =====================================================================
//  universal service structures
// =====================================================================

type Filter struct {
	UserID   int
	Cursor   int
	Offset   int
	SortType int
}

// =====================================================================
//  user Query && Command struct
// =====================================================================

type userQueryUseCase struct {
	queryRepo       UserQueryRepository
	sessionProvider SessionProvider
}

type userCommandUseCase struct {
	commandRepo     UserCommandRepository
	sessionProvider SessionProvider
}

// =====================================================================
//  user Query && Command service structures
// =====================================================================

type CreateReport struct {
	UserID       int
	UserMessage  string
	UserReason   int
	ContentID    int
	ContentSubID string
	ContentType  int
}

// =====================================================================
//  moderator Query && Command struct
// =====================================================================

type moderatorQueryUseCase struct {
	queryRepo       ModeratorQueryRepository
	sessionProvider SessionProvider
}

type moderatorCommandUseCase struct {
	commandRepo     ModeratorCommandRepository
	queryRepo       ModeratorQueryRepository
	sessionProvider SessionProvider
}

// =====================================================================
//  moderator Query && Command service structures
// =====================================================================

type CommandRequest struct {
	ModeratorID int
	CaseID      int
	Message     string
}

type TransferCase struct {
	OldModeratorID int
	NewModeratorID int
	CaseID         int
	Message        string
}
