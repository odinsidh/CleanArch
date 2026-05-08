package entity

import (
	"time"
)

// =====================================================================
//  Reaction
// =====================================================================

type Reaction struct {
	ReactionID   int
	ReactionType ReactionType
	UserID       UserID
	TargetID     TargetID
	TargetType   TargetType
	Timestamp    time.Time
}

func NewReaction(ReactionType ReactionType, UserID int, TargetID int, TargetType string) (*Reaction, error) {
	var err error

	// rt, err := NewReactionType(ReactionType)
	// if err != nil {
	// 	return nil, err
	// }

	uID, err := NewUserID(UserID)
	if err != nil {
		return nil, err
	}

	tID, err := NewTargetID(TargetID)
	if err != nil {
		return nil, err
	}

	tt, err := NewTargetType(TargetType)
	if err != nil {
		return nil, err
	}

	output := &Reaction{
		ReactionType: ReactionType,
		UserID:       uID,
		TargetID:     tID,
		TargetType:   tt,
		Timestamp:    time.Now(),
	}

	return output, nil
}

func (self *Reaction) TouchPositive() {
	if self.ReactionType == Positive {
		self.ReactionType = None
		return
	}

	self.ReactionType = Positive
}

func (self *Reaction) TouchNegative() {
	if self.ReactionType == Negative {
		self.ReactionType = None
		return
	}

	self.ReactionType = Negative
}
