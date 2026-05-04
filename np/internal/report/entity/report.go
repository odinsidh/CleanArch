package entity

import (
	"errors"
	"time"
)

type Report struct {
	CaseID      int
	ReportID    int
	UserID      int
	UserMessage UserMessage
	Reason      Reason
	CreatedAt   time.Time
}

var (
	ErrUserIDEqualToZero error = errors.New("userID was equal to zero")
)

func NewReport(userID int, userMessage string, reason int) (*Report, error) {
	if userID == 0 {
		return nil, ErrUserIDEqualToZero
	}

	reportCreatedAt := time.Now()

	reportMessage, err := NewUserMessage(userMessage)
	if err != nil {
		return nil, err
	}

	reportReason, err := NewReason(reason)
	if err != nil {
		return nil, err
	}

	output := Report{
		UserID:      userID,
		UserMessage: *reportMessage,
		Reason:      *reportReason,
		CreatedAt:   reportCreatedAt,
	}

	return &output, nil
}

func (self *Report) AddCaseID(input int) {
	self.CaseID = input
}

func (self *Report) AddReportID(input int) {
	self.ReportID = input
}
