package usecase

import (
	"context"
	"fmt"
	"newsportal/internal/article/dto"
	eArticle "newsportal/internal/article/entity"
	eAuthorization "newsportal/internal/authorization/entity"
)

func (self *commandUseCase) CreateArticle(ctx context.Context, authorID int, articleType string, title string, content string) (*dto.ArticleItem, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	op := " | article > commandUseCase > CreateArticle | "
	input := fmt.Sprintf("input > authorID: [%v] articleType: [%v] title: [%v] content: [%v]", authorID, articleType, title, content)

	// =====================================================================
	//  get user session
	// =====================================================================
	session, err := self.sessionProvider.Session(ctx, authorID)
	if err != nil {
		return nil, fmt.Errorf("%w %v %v",
			err, op, input)
	}

	// =====================================================================
	//  read permission
	// =====================================================================
	if !session.Can(eAuthorization.Article, eAuthorization.CanCreate) {
		return nil, fmt.Errorf("%w %v %v",
			err, op, input,
		)
	}

	// =====================================================================
	//  create article entity
	// =====================================================================
	newArticle, err := eArticle.NewArticle(authorID, articleType, title, content)
	if err != nil {
		return nil, err
	}

	// =====================================================================
	//  create article entity
	// =====================================================================
	createdArticle, err := self.commandRepository.CreateArticle(ctx, newArticle)
	if err != nil {
		return nil, err
	}

	// =====================================================================
	//  thread initialization
	// =====================================================================
	err = self.threadInitialize.Initialize(ctx, createdArticle.ArticleType, createdArticle.ArticleID)
	if err != nil {
		_ = self.commandRepository.HardDelete(ctx, createdArticle)
		// тут скипаем ошибку, когда у нас все отвалилось (моргает бд с комментами, пока не решаем проблему)
		return nil, err
	}

	// =====================================================================
	//  prepare dto
	// =====================================================================
	articleActions := dto.ArticleActions{
		CanRead:   true,
		CanModify: true,
		CanDelete: true,
	}

	createdArticleDTO := dto.NewArticleItem(createdArticle.ArticleID,
		createdArticle.ArticleType,
		string(createdArticle.Content),
		createdArticle.ArticleType,
		createdArticle.AuthorID,
		createdArticle.CreatedAT,
		articleActions)

	return createdArticleDTO, nil
}

func (self *commandUseCase) ModifyArticle(ctx context.Context, userID int, articleType string, articleID int, title string, content string) (*dto.ArticleItem, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	op := " | article > commandUseCase > ModifyArticle | "
	input := fmt.Sprintf("input > authorID: [%v] articleType: [%v] articleID: [%v] title: [%v] content: [%v]", userID, articleType, articleID, title, content)

	// =====================================================================
	//  get user session
	// =====================================================================
	session, err := self.sessionProvider.Session(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w %v %v",
			err, op, input)
	}

	// =====================================================================
	//  read permission
	// =====================================================================
	permissionDomain := eAuthorization.Article
	permissionRead := eAuthorization.CanRead
	if !session.Can(permissionDomain, permissionRead) {
		return nil, fmt.Errorf("%w %v %v domain: [%v] permission: [%v]",
			ErrUserDontHavePermission, op, input, permissionDomain, permissionRead)
	}

	// =====================================================================
	//  get article from repository
	// =====================================================================
	readedArticle, err := self.commandRepository.ReadArticle(ctx, articleType, articleID)
	if err != nil {
		return nil, fmt.Errorf("%w %v %v",
			err, op, input,
		)
	}

	// =====================================================================
	//  ask entity about permission
	// =====================================================================
	permissionModify := eAuthorization.CanModify
	if !readedArticle.CanBeUpdatedByID(userID, session.Can(permissionDomain, permissionModify)) {
		return nil, fmt.Errorf("%w %v %v domain: [%v] permission: [%v]",
			ErrUserDontHavePermission, op, input, permissionDomain, permissionModify)

	}

	// =====================================================================
	//  update article entity
	// =====================================================================
	err = readedArticle.UpdateArticle(title, content)
	if err != nil {
		return nil, fmt.Errorf("%w %v %v",
			err, op, input,
		)

	}

	// =====================================================================
	//  update article in repository
	// =====================================================================
	updatedArticle, err := self.commandRepository.ModifyArticle(ctx, readedArticle)
	if err != nil {
		return nil, err
	}

	// =====================================================================
	//  dto
	// =====================================================================
	permissionDelete := eAuthorization.CanDelete
	articleActions := dto.ArticleActions{
		CanRead:   true,
		CanModify: true,
		CanDelete: readedArticle.CanBeUpdatedByID(userID, session.Can(permissionDomain, permissionDelete)),
	}

	updatedArticleDTO := dto.NewArticleItem(updatedArticle.ArticleID,
		updatedArticle.ArticleType,
		string(updatedArticle.Content),
		updatedArticle.ArticleType,
		updatedArticle.AuthorID,
		updatedArticle.CreatedAT,
		articleActions)

	return updatedArticleDTO, nil
}

func (self *commandUseCase) DeleteArticle(ctx context.Context, userID int, articleType string, articleID int) error {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	op := " | article > commandUseCase > DeleteArticle | "
	input := fmt.Sprintf("input > userID: [%v] articleType: [%v] articleID: [%v] ", userID, articleType, articleID)

	// =====================================================================
	//  get user session
	// =====================================================================
	session, err := self.sessionProvider.Session(ctx, userID)
	if err != nil {
		return fmt.Errorf("%w %v %v",
			err, op, input)
	}

	// =====================================================================
	//  read permission
	// =====================================================================
	permissionDomain := eAuthorization.Article
	permissionRead := eAuthorization.CanRead
	if !session.Can(permissionDomain, permissionRead) {
		return fmt.Errorf("%w %v %v domain: [%v] permission: [%v]",
			ErrUserDontHavePermission, op, input, permissionDomain, permissionRead)
	}

	// =====================================================================
	//  get article from repository
	// =====================================================================
	readedArticle, err := self.commandRepository.ReadArticle(ctx, articleType, articleID)
	if err != nil {
		return fmt.Errorf("%w %v %v",
			err, op, input,
		)
	}

	// =====================================================================
	//  ask entity about permission
	// =====================================================================
	permissionDelete := eAuthorization.CanDelete
	if !readedArticle.CanBeDeletedByID(userID, session.Can(permissionDomain, permissionDelete)) {
		return fmt.Errorf("%w %v %v domain: [%v] permission: [%v]",
			ErrUserDontHavePermission, op, input, permissionDomain, permissionDelete)
	}

	// =====================================================================
	//  soft delete from repository
	// =====================================================================
	readedArticle.DeleteArticle()

	err = self.commandRepository.DeleteArticle(ctx, readedArticle)
	if err != nil {
		return fmt.Errorf("%w %v %v",
			err, op, input,
		)
	}

	return nil
}

func (self *commandUseCase) PublishArticle(ctx context.Context, userID int, articleType string, articleID int) error {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	op := " | article > commandUseCase > PublishArticle | "
	input := fmt.Sprintf("input > authorID: [%v] articleType: [%v] articleID: [%v] ", userID, articleType, articleID)

	// =====================================================================
	//  get user session
	// =====================================================================
	session, err := self.sessionProvider.Session(ctx, userID)
	if err != nil {
		return fmt.Errorf("%w %v %v",
			err, op, input)
	}

	// =====================================================================
	//  read permission
	// =====================================================================
	if !session.Can(eAuthorization.Article, eAuthorization.CanRead) {
		return fmt.Errorf("%w %v %v",
			err, op, input,
		)
	}

	// =====================================================================
	//  get article from repository
	// =====================================================================
	readedArticle, err := self.commandRepository.ReadArticle(ctx, articleType, articleID)
	if err != nil {
		return fmt.Errorf("%w %v %v",
			err, op, input,
		)
	}

	// =====================================================================
	//  ask entity about permission
	// =====================================================================
	permissionDomain := eAuthorization.Article
	permissionModify := eAuthorization.CanModify
	if !readedArticle.CanBePublishByID(userID, session.Can(permissionDomain, permissionModify)) {
		return fmt.Errorf("%w %v %v domain: [%v] permission: [%v]",
			ErrUserDontHavePermission, op, input, permissionDomain, permissionModify)
	}

	// =====================================================================
	//  article change status to published in repository
	// =====================================================================
	err = self.commandRepository.PublishArticle(ctx, readedArticle)
	if err != nil {
		return fmt.Errorf("%w %v %v",
			err, op, input,
		)
	}

	return nil
}
