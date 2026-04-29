package domain

type PostStatus string

const (
	PostStatusDraft     PostStatus = "draft"
	PostStatusPublished PostStatus = "published"
	PostStatusArchived  PostStatus = "archived"
)

type Post struct {
	ID          int        `db:"id" json:"id"`
	Title       string     `db:"title" json:"title"`
	Slug        string     `db:"slug" json:"slug"`
	Content     string     `db:"content" json:"content"`
	Summary     string     `db:"summary" json:"summary"`
	Thumbnail   string     `db:"thumbnail" json:"thumbnail,omitempty"`
	AuthorID    int        `db:"author_id" json:"authorId"`
	CategoryID  *int       `db:"category_id" json:"categoryId,omitempty"`
	Status      PostStatus `db:"status" json:"status"`
	ViewCount   int        `db:"view_count" json:"viewCount"`
	PublishedAt *int64     `db:"published_at" json:"publishedAt,omitempty"`
	CreatedAt   int64      `db:"created_at" json:"createdAt"`
	UpdatedAt   int64      `db:"updated_at" json:"updatedAt"`
}
