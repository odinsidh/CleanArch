package usecase

import articleDTO "newsportal/internal/article/dto"

// =====================================================================
//  Query UseCase helpers
// =====================================================================

func extractAuthorID(articleItems []articleDTO.ArticleItem) []int {
	var container map[int]struct{} = make(map[int]struct{}, len(articleItems))
	var output []int = make([]int, 0, len(articleItems))
	for index := range articleItems {
		authorID := articleItems[index].AuthorID
		if _, ok := container[authorID]; !ok {
			container[authorID] = struct{}{}
			output = append(output, authorID)
		}
	}

	return output
}

func extractArticleID(articleItems []articleDTO.ArticleItem) []int {
	var container map[int]struct{} = make(map[int]struct{}, len(articleItems))
	var output []int = make([]int, 0, len(articleItems))
	for index := range articleItems {
		articleID := articleItems[index].ArticleID
		if _, ok := container[articleID]; !ok {
			container[articleID] = struct{}{}
			output = append(output, articleID)
		}
	}

	return output
}
