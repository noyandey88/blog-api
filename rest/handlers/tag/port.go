package tag

import "github.com/noyandey88/blog-api/domain"

type Service interface {
	Create(tag domain.Tag) (*domain.Tag, error)
	List() ([]*domain.Tag, error)
	Get(id int) (*domain.Tag, error)
	Update(tag domain.Tag) (*domain.Tag, error)
	Delete(id int) error
}
