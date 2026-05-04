package entity

import (
	"errors"
	"time"
)

type Case struct {
	CaseID       int
	ContentID    int
	ContentSubID string
	ContentType  ContentType
	ModeratorID  int
	Status       Status
	Reports      []Report
	History      []History
	CreatedAT    time.Time
}

var (
	reattachModeratorMessage string = "case moderator was changed"
)

var (
	ErrContentSubIdIsRequired error = errors.New("contentSubID is required")
	// ErrStatusAlreadyApplied               error = errors.New("case status already applied")
	ErrStatusClosedCase         error = errors.New("case was already closed")
	ErrCaseIDEqualToZero        error = errors.New("case id is equal to zero")
	ErrCaseNotHaveNewStatus     error = errors.New("case not have new status")
	ErrUserWasAlreadyHaveReport error = errors.New("this user was already create report to this content")
	// ErrModeratorWasAlreadyAttach          error = errors.New("in this case moderator already attach")
	ErrModeratorMessageTooLong error = errors.New("moderator message was too long")
	// ErrCaseModeratorNotEqualToCurrent     error = errors.New("moderator try to modify case, from another admin")
	ErrOldModeratorAndNewModeratorIsEqual error = errors.New("moderator try to reattach case to self")
	ErrCurrentModeratorAlreadyLeadCase    error = errors.New("current moderator already lead that case")
	ErrAnotherModeratorAlreadyLeadCase    error = errors.New("another moderator already lead that case")
)

func NewCase(contentID int, contentSubID string, contentType int) (*Case, error) {
	if contentID == int(Article) && contentSubID == "" {
		return nil, ErrContentSubIdIsRequired
	}

	caseContentType, err := NewContentType(contentType)
	if err != nil {
		return nil, err
	}

	caseNewStatus, err := NewStatus(int(New))
	if err != nil {
		return nil, err
	}

	caseCreatedAt := time.Now()

	output := Case{
		ContentID:    contentID,
		ContentSubID: contentSubID,
		ContentType:  *caseContentType,
		Status:       *caseNewStatus,
		CreatedAT:    caseCreatedAt,
	}

	return &output, nil
}

func (self *Case) IsClosed() error {
	if self.Status == Banned || self.Status == Disagree {
		return ErrStatusClosedCase
	}
	return nil
}

// Метод исключительно для базы данных, прогрузка Case структуры, не использоватьв в других целях
func (self *Case) AddReports(input ...Report) {
	self.Reports = append(self.Reports, input...)
}

// Метод исключительно для базы данных, прогрузка Case структуры, не использоватьв в других целях
func (self *Case) AddHistory(input ...History) {
	self.History = append(self.History, input...)
}

func (self *Case) AddReport(userID int, userMessage string, reason int) error {
	err := self.IsClosed()
	if err != nil {
		return err
	}

	for index := range self.Reports {
		if self.Reports[index].UserID == userID {
			return ErrUserWasAlreadyHaveReport
		}
	}

	outputReport, err := NewReport(userID, userMessage, reason)
	if err != nil {
		return err
	}

	self.Reports = append(self.Reports, *outputReport)
	return nil
}

func (self *Case) TakeInWork(moderatorID int, moderatorMessage string) error {
	if self.Status != New {
		return ErrCaseNotHaveNewStatus
	}

	if self.CaseID == 0 {
		return ErrCaseIDEqualToZero
	}

	if self.ModeratorID != 0 && self.ModeratorID != moderatorID {
		return ErrAnotherModeratorAlreadyLeadCase
	}

	if self.ModeratorID != 0 && self.ModeratorID == moderatorID {
		return ErrCurrentModeratorAlreadyLeadCase
	}

	self.ModeratorID = moderatorID
	self.Status = InWork
	singleHistory, err := NewHistory(self.CaseID, moderatorID, moderatorMessage, int(self.Status))
	if err != nil {
		return err
	}
	self.History = append(self.History, *singleHistory)

	return nil
}

func (self *Case) AddModeratorMessage(moderatorID int, moderatorMessage string) error {
	err := self.IsClosed()
	if err != nil {
		return err
	}

	if self.CaseID == 0 {
		return ErrCaseIDEqualToZero
	}

	if self.ModeratorID != moderatorID {
		return ErrAnotherModeratorAlreadyLeadCase
	}

	if len(moderatorMessage) > 1000 {
		return ErrModeratorMessageTooLong
	}

	historyUnit, err := NewHistory(self.CaseID, moderatorID, moderatorMessage, int(self.Status))
	if err != nil {
		return err
	}
	self.History = append(self.History, *historyUnit)

	return nil

}

func (self *Case) AddModeratorDecisionResolve(moderatorID int, moderatorMessage string) error {
	err := self.IsClosed()
	if err != nil {
		return err
	}

	if self.CaseID == 0 {
		return ErrCaseIDEqualToZero
	}

	if self.ModeratorID != moderatorID {
		return ErrAnotherModeratorAlreadyLeadCase
	}

	if len(moderatorMessage) > 1000 {
		return ErrModeratorMessageTooLong
	}

	self.Status = Banned
	caseHistory, err := NewHistory(self.CaseID, moderatorID, moderatorMessage, int(self.Status))
	if err != nil {
		return err
	}
	self.AddHistory(*caseHistory)

	return nil
}

func (self *Case) AddModeratorDecisionReject(moderatorID int, moderatorMessage string) error {
	err := self.IsClosed()
	if err != nil {
		return err
	}

	if self.CaseID == 0 {
		return ErrCaseIDEqualToZero
	}

	if self.ModeratorID != moderatorID {
		return ErrAnotherModeratorAlreadyLeadCase
	}

	if len(moderatorMessage) > 1000 {
		return ErrModeratorMessageTooLong
	}

	self.Status = Disagree
	caseHistory, err := NewHistory(self.CaseID, moderatorID, moderatorMessage, int(self.Status))
	if err != nil {
		return err
	}
	self.AddHistory(*caseHistory)

	return nil
}

func (self *Case) TransferCaseToAnotherModerator(moderatorID, newModeratorID int, message string) error {
	err := self.IsClosed()
	if err != nil {
		return err
	}

	if self.CaseID == 0 {
		return ErrCaseIDEqualToZero
	}

	if self.ModeratorID != moderatorID {
		return ErrAnotherModeratorAlreadyLeadCase
	}

	if moderatorID == newModeratorID {
		return ErrOldModeratorAndNewModeratorIsEqual
	}

	if len(message) > 0 {
		err := self.AddModeratorMessage(moderatorID, message)
		if err != nil {
			return err
		}
	} else {
		err = self.AddModeratorMessage(moderatorID, reattachModeratorMessage)
		if err != nil {
			return err
		}
	}
	self.ModeratorID = newModeratorID

	return nil
}
