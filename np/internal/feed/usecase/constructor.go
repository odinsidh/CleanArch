package usecase

// =====================================================================
//  Query Usecase constructor
// =====================================================================

func NewFeed(sessionProvider SessionProvider, articleQuery ArticleQuery, userQuery UserQuery, commentsQuery CommentsQuery) Feed {
	return &feed{
		sessionProvider: sessionProvider,
		articleQuery:    articleQuery,
		userQuery:       userQuery,
		commentsQuery:   commentsQuery,
	}
}
