package tag

import (
	"github.com/noyandey88/blog-api/domain"
	tagHandler "github.com/noyandey88/blog-api/rest/handlers/tag"
)

type Service interface {
	tagHandler.Service
}

type TagRepo interface {
	Create(tag domain.Tag) (*domain.Tag, error)
	List() ([]*domain.Tag, error)
	Get(id int) (*domain.Tag, error)
	Update(tag domain.Tag) (*domain.Tag, error)
	Delete(id int) error
}
