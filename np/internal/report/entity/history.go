package entity

import (
	"errors"
	"time"
)

type History struct {
	CaseID           int
	HistoryID        int
	ModeratorID      int
	ModeratorMessage string
	Status           Status
	CreatedAt        time.Time
}

var (
	ErrHistoryCaseIDEqualToZero      error = errors.New("history case id is equal to zero")
	ErrHistoryModeratorIDEqualToZero error = errors.New("history moderator id is equal to zero")
	ErrHistoryMessageTooLong         error = errors.New("history message was too long")
)

func NewHistory(caseID int, moderatorID int, moderatorMessage string, status int) (*History, error) {
	// вот тут вопросики, выносить это в отдельный класс как будто бы не нужно,
	// но при этом мусорить логикой проверки тоже под вопросом
	if caseID == 0 {
		return nil, ErrHistoryCaseIDEqualToZero
	}

	if moderatorID == 0 {
		return nil, ErrHistoryModeratorIDEqualToZero
	}

	if len(moderatorMessage) > 1000 {
		return nil, ErrHistoryMessageTooLong
	}

	historyStatus, err := NewStatus(status)
	if err != nil {
		return nil, err
	}

	historyCreatedAt := time.Now()

	output := History{
		CaseID:           caseID,
		ModeratorID:      moderatorID,
		ModeratorMessage: moderatorMessage,
		Status:           *historyStatus,
		CreatedAt:        historyCreatedAt,
	}

	return &output, nil
}
