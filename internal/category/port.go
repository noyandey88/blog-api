package category

import (
	"github.com/noyandey88/blog-api/domain"
	categoryHandler "github.com/noyandey88/blog-api/rest/handlers/category"
)

type Service interface {
	categoryHandler.Service
}

type CategoryRepo interface {
	Create(category domain.Category) (*domain.Category, error)
	List(page, limit int64) ([]*domain.Category, error)
	Get(id int) (*domain.Category, error)
	Update(category domain.Category) (*domain.Category, error)
	Delete(id int) error
}
