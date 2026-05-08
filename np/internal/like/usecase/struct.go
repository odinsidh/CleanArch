package usecase

import (
	eLike "newsportal/internal/like/entity"
)

// =====================================================================
//  universal service structures
// =====================================================================

type Deps struct {
	SessionProvider   SessionProvider
	QueryRepo         QueryRepository
	CommandRepository CommandRepository
}

// =====================================================================
//  Query struct
// =====================================================================

type CountRequest struct {
	Target
}

type CountResponse struct {
	Target
	Count map[eLike.ReactionType]int
}

type Target struct {
	ID   eLike.TargetID
	Type eLike.TargetType
}

type queryUseCase struct {
	sessionProvider SessionProvider
	queryRepo       QueryRepository
}

// =====================================================================
//  Command struct
// =====================================================================

type commandUseCase struct {
	sessionProvider   SessionProvider
	commandRepository CommandRepository
}
