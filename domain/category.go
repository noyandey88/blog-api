package domain

type Category struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Slug        string `json:"slug" db:"slug"`
	Description string `json:"description" db:"description"`
	CreatedAt   int64  `json:"createdAt" db:"created_at"`
	UpdatedAt   int64  `json:"updatedAt" db:"updated_at"`
}
