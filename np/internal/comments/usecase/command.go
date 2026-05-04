package usecase

import (
	"context"
	"fmt"
	eAuthorization "newsportal/internal/authorization/entity"
	eComments "newsportal/internal/comments/entity"
)

func (self *commandUseCase) AddComment(ctx context.Context, userID int, selector Selector, replyToID int, value string) error {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "comments > commandUseCase > AddComment"
	input := fmt.Sprintf("userID: [%v] selector: [%v] replyToID: [%v] value: [%v]", userID, selector, replyToID, value)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if userID == 0 {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}
	if selector.TargetID == 0 {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrTargetIDIsEqualToZero, input,
		)
	}
	if selector.TargetType == "" {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrTargetTypeNotFound, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Messages
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead, eAuthorization.CanCreate}
	_, err := havePermission(ctx, self.sessionProvider, userID, permissionDomain, permissionToCheck...)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	comment, err := eComments.NewComment(selector.TargetType, selector.TargetID, replyToID, userID, value)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	if comment.ReplyMessageID != 0 {
		replyComment, err := self.commandRepo.GetCommentByID(ctx, string(comment.TargetType), comment.TargetID, comment.ReplyMessageID)

		if err != nil {
			return fmt.Errorf("%s : %w context(%v)",
				operation, ErrReplyCommentNotFound, input)
		}

		if comment.TargetID != replyComment.TargetID || comment.TargetType != replyComment.TargetType {
			return fmt.Errorf("%s : %w context(%v)",
				operation, ErrTryingToAddCommentWithNotEqualThreadIDandThreadType, input)
		}
	}

	err = self.commandRepo.AddComment(ctx, comment)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return nil
}

func (self *commandUseCase) UpdateComment(ctx context.Context, userID int, selector Selector, messageID int, value string) error {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "comments > commandUseCase > UpdateComment"
	input := fmt.Sprintf("userID: [%v] selector: [%v] messageID: [%v] value: [%v]", userID, selector, messageID, value)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if userID == 0 {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}
	if selector.TargetID == 0 {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrTargetIDIsEqualToZero, input,
		)
	}
	if selector.TargetType == "" {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrTargetTypeNotFound, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Messages
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead}
	session, err := havePermission(ctx, self.sessionProvider, userID, permissionDomain, permissionToCheck...)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	comment, err := self.commandRepo.GetCommentByID(ctx, selector.TargetType, selector.TargetID, messageID)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	permissionModify := eAuthorization.CanModify
	err = comment.Update(userID, value, session.Can(permissionDomain, permissionModify))
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	err = self.commandRepo.UpdateComment(ctx, comment)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return nil
}

func (self *commandUseCase) DeleteComment(ctx context.Context, userID int, selector Selector, messageID int) error {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "comments > commandUseCase > DeleteComment"
	input := fmt.Sprintf("userID: [%v] selector: [%v] messageID: [%v]", userID, selector, messageID)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if userID == 0 {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrUserIDIsEqualToZero, input,
		)
	}
	if selector.TargetID == 0 {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrTargetIDIsEqualToZero, input,
		)
	}
	if selector.TargetType == "" {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrTargetTypeNotFound, input,
		)
	}
	if messageID == 0 {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrMessageIDIsEqualToZero, input,
		)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	permissionDomain := eAuthorization.Messages
	permissionToCheck := []eAuthorization.Permission{eAuthorization.CanRead}
	session, err := havePermission(ctx, self.sessionProvider, userID, permissionDomain, permissionToCheck...)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	comment, err := self.commandRepo.GetCommentByID(ctx, selector.TargetType, selector.TargetID, messageID)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	permissionDelete := eAuthorization.CanDelete
	err = comment.Deactivate(userID, session.Can(permissionDomain, permissionDelete))
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	err = self.commandRepo.DeleteComment(ctx, comment)
	if err != nil {
		return fmt.Errorf("%s : %w context(%v)",
			operation, err, input)
	}

	return nil
}

func (self *commandUseCase) Initialize(ctx context.Context, selector Selector) error {
	// TODO, тут наверное разве что, можем вынести проверку на доступные threadType ,
	// и можем ли мы с ними взаимодействовать, чтобы не давать задание базе, на работу с несуществующей базой
	return self.commandRepo.CreateThread(ctx, selector)
}
