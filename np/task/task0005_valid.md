Задача:
Сделать кормушку \ главная страница \ каталог
feed

Структура
type Feed struct {
    FeedItems []FeedItem
    Page int
    Pagination int
}

type FeedItem struct {
    ArticleID int
    ArticleTitle string
    ArticlePreview string
    ArticleAuthorUsername int
    PublishedAt time.Time
    CommentsCount int
}

const PaginationMinValue int = 10
const PaginationMaxValue int = 20
const ArticlePreviewLenght int = 100

Нужны облегченные интерфейсы (+ query)
User 
> получить Username по ID

Article
> Получить X статей, по Y категории, отправная точка Z, метод сортировки ASC|DESC
ONLY STATUS = PUBLISHED

Comment
> Получить количество комментариев к статье



Что отдаем > Склеенное DTO





Вводные:
Получаем только опубликованные статьи (Status = Published)
Тип сортировки (сначала свежие, сначала старые)
Пагинация (по 10-20 штук за раз)
Каждая статья выводится не полностью, с неё забирается только preview

Показываем пользователю (Заголовок, автор, дата, кусочек текста (превью))
Количество комментариев к статье


