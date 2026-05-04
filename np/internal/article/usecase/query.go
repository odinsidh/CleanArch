package usecase

import (
	"context"
	"fmt"
	"newsportal/internal/article/dto"
	"newsportal/internal/authorization/entity"
	eAuthorization "newsportal/internal/authorization/entity"
)

func (self *queryUseCase) GetArticles(ctx context.Context, userID int, articleType string, cursor int, limit int, sortType int) ([]dto.ArticleItem, error) {
	// TODO
	// CURSOR >= 0 // LIMIT >= && <= 20 (set as constant) // redis cache will be living her

	// =====================================================================
	//  Error tracing
	// =====================================================================
	op := " | article > queryUseCase > GetArticles | "
	input := fmt.Sprintf("input > userID: [%v] articleType: [%v] cursor: [%v] limit: [%v] sortType: [%v]",
		userID, articleType, cursor, limit, sortType)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if articleType == "" {
		return nil, fmt.Errorf("%w %v %v",
			ErrArticleTypeNotFound, op, input)
	}

	if sortType == 0 {
		return nil, fmt.Errorf("%w %v %v",
			ErrSortTypeIsEqualToZero, op, input)
	}

	if userID == 0 {
		return nil, fmt.Errorf("%w %v %v",
			ErrUserIDIsEqualToZero, op, input)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	session, err := self.sessionProvider.Session(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w %v %v",
			err, op, input)
	}

	permissionDomain := eAuthorization.Article
	permissionRead := eAuthorization.CanRead
	if !session.Can(permissionDomain, permissionRead) {
		return nil, fmt.Errorf("%w %v %v domain: [%v] permission: [%v]",
			ErrUserDontHavePermission, op, input, permissionDomain, permissionRead)
	}

	// =====================================================================
	//  get articles from repository
	// =====================================================================
	articles, err := self.queryRepository.GetArticles(ctx, articleType, cursor, limit, sortType)
	if err != nil {
		return nil, fmt.Errorf("%w %v %v",
			err, op, input)
	}

	// =====================================================================
	//  permission acting on dto
	// =====================================================================
	permissionModify := session.Can(entity.Article, entity.CanModify)
	permissionDelete := session.Can(entity.Article, entity.CanDelete)

	// =====================================================================
	//  attach permissions to dto
	// =====================================================================
	for index := range articles {
		articles[index].ArticleActions.CanRead = true

		isAuthor := articles[index].AuthorID == userID
		articles[index].ArticleActions.CanModify = permissionModify || isAuthor
		articles[index].ArticleActions.CanDelete = permissionDelete || isAuthor
	}

	return articles, nil
}

func (self *queryUseCase) GetArticle(ctx context.Context, userID int, articleType string, articleID int) (*dto.ArticleItem, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	op := " | article > queryUseCase > GetArticle | "
	input := fmt.Sprintf("input > userID: [%v] articleType: [%v] articleID: [%v]", userID, articleType, articleID)

	// =====================================================================
	//  get user session
	// =====================================================================
	session, err := self.sessionProvider.Session(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w %v %v",
			err, op, input)
	}

	permissionDomain := eAuthorization.Article
	permissionRead := eAuthorization.CanRead
	if !session.Can(permissionDomain, permissionRead) {
		return nil, fmt.Errorf("%w %v %v domain: [%v] permission: [%v]",
			ErrUserDontHavePermission, op, input, permissionDomain, permissionRead)
	}

	// =====================================================================
	//  get article from repository
	// =====================================================================
	readedArticle, err := self.queryRepository.GetArticle(ctx, articleType, articleID)
	if err != nil {
		return nil, err
	}

	// =====================================================================
	//  attach permission
	// =====================================================================
	permissionModify := eAuthorization.CanModify
	permissionDetele := eAuthorization.CanDelete

	isAuthor := readedArticle.AuthorID == userID
	articleActions := dto.ArticleActions{
		CanRead:   true,
		CanModify: session.Can(permissionDomain, permissionModify) || isAuthor,
		CanDelete: session.Can(permissionDomain, permissionDetele) || isAuthor,
	}

	// =====================================================================
	//  dto
	// =====================================================================
	readedArticleDTO := dto.NewArticleItem(readedArticle.ArticleID,
		readedArticle.ArticleType,
		string(readedArticle.Content),
		readedArticle.ArticleType,
		readedArticle.AuthorID,
		readedArticle.CreatedAT,
		articleActions)

	return readedArticleDTO, nil
}
