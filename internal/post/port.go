package post

import (
	"github.com/noyandey88/blog-api/domain"
	postHandler "github.com/noyandey88/blog-api/rest/handlers/post"
)

type Service interface {
	postHandler.Service
}

type PostRepo interface {
	Create(post domain.Post) (*domain.Post, error)
	List(page, limit int64) ([]*domain.Post, error)
	Count() (int64, error)
	Get(id int) (*domain.Post, error)
	Update(post domain.Post) (*domain.Post, error)
	Delete(id int) error
}
