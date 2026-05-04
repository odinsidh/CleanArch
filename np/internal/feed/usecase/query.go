package usecase

import (
	"context"
	"fmt"
	eAuthorization "newsportal/internal/authorization/entity"
	feedDTO "newsportal/internal/feed/dto"
	eFeed "newsportal/internal/feed/entity"
)

func (self *feed) GetFeed(ctx context.Context, filter Filter) (*feedDTO.FeedDTO, error) {
	// =====================================================================
	//  Error tracing
	// =====================================================================
	operation := "feed > queryUseCase > GetFeed"
	input := fmt.Sprintf("input > userID: [%v] articleType: [%v] cursor: [%v] limit: [%v] sortType: [%v]",
		filter.UserID, filter.ArticleType, filter.Cursor, filter.Limit, filter.SortType)

	// =====================================================================
	//  Use case checkers
	// =====================================================================
	if filter.UserID == 0 {
		return nil, fmt.Errorf("%s : %w context:(%v)",
			operation, ErrUserIDIsEqualToZero, input)
	}

	if filter.ArticleType == "" {
		return nil, fmt.Errorf("%s : %w context:(%v)",
			operation, ErrArticleTypeNotFound, input)
	}

	_, err := eFeed.NewSortType(filter.SortType)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context:(%v)",
			operation, err, input)
	}

	// =====================================================================
	//  get user session
	// =====================================================================
	session, err := self.sessionProvider.Session(ctx, filter.UserID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context:(%v)",
			operation, err, input)
	}

	permissionDomain := eAuthorization.Feed
	permissionRead := eAuthorization.CanRead
	if !session.Can(permissionDomain, permissionRead) {
		return nil, fmt.Errorf("%s : %w context:(%v) domain: [%v] permission: [%v]",
			operation, ErrUserDontHavePermission, input, permissionDomain, permissionRead)
	}

	// =====================================================================
	//  Use case core logick
	// =====================================================================
	// TODO limit as value object with checks on minimal x maniximal value?
	// if cursor equal to zero = take first, based on sortType (headache of repo side)
	articles, err := self.articleQuery.GetArticles(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context:(%v)",
			operation, err, input,
		)
	}

	// TODO тут на подумать, а действительно ли пользователю не отдавать статьи, если что-то случилось на стороне получения пользователей
	authorIDs := extractAuthorID(articles)
	usernamesByID, err := self.userQuery.GetUsernamesByID(ctx, authorIDs...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context:(%v)",
			operation, err, input,
		)
	}

	// TODO тут на подумать, а действительно ли пользователю не отдавать статьи, если что-то случилось на стороне получения комментариев
	articleIDs := extractArticleID(articles)
	commentCountByID, err := self.commentsQuery.GetCommentsCountByID(ctx, filter.ArticleType, articleIDs...)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context:(%v)",
			operation, err, input,
		)
	}

	// =====================================================================
	//  dto
	// =====================================================================
	output := &feedDTO.FeedDTO{
		FeedItems: make([]feedDTO.FeedItemDTO, 0, len(articles)),
	}
	for _, value := range articles {
		output.FeedItems = append(output.FeedItems, feedDTO.FeedItemDTO{
			ArticleID:      value.ArticleID,
			ArticleTitle:   value.ArticleTitle,
			ArticleContent: value.ArticleContent,
			ArticleType:    value.ArticleType,
			AuthorID:       value.AuthorID,
			Timestamp:      value.PublishedAt,
			ArticleActions: feedDTO.ArticleActions(value.ArticleActions),
		})
	}

	err = output.AttachUsername(usernamesByID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context:(%v)",
			operation, err, input,
		)
	}

	err = output.AttachCommentAmount(commentCountByID)
	if err != nil {
		return nil, fmt.Errorf("%s : %w context:(%v)",
			operation, err, input,
		)
	}

	return output, nil
}
