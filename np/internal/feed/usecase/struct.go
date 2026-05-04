package usecase

// =====================================================================
//  Query UseCase struct
// =====================================================================

type feed struct {
	sessionProvider SessionProvider
	articleQuery    ArticleQuery
	userQuery       UserQuery
	commentsQuery   CommentsQuery
}

type Filter struct {
	UserID      int
	ArticleType string
	Cursor      int
	Limit       int
	SortType    int
}
