package category

import "github.com/noyandey88/blog-api/domain"

type Service interface {
	Create(category domain.Category) (*domain.Category, error)
	FindAll() (*[]domain.Category, error)
	FindById(id int) (*domain.Category, error)
	Update(category domain.Category) (*domain.Category, error)
	Delete(id int) error
}
