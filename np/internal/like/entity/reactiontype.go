package entity

// =====================================================================
//  ReactionType
// =====================================================================

type ReactionType int

const (
	None ReactionType = iota + 1
	Positive
	Negative
)

func NewReactionType(selectedReaction int) (ReactionType, error) {
	rt := ReactionType(selectedReaction)

	switch rt {
	case Positive, Negative:
		return rt, nil

	default:
		return 0, ErrReactionTypeNotExist
	}
}
