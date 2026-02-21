package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/noyandey88/blog-api/domain"
	"github.com/noyandey88/blog-api/internal/category"
)

type CategoryRepo interface {
	category.CategoryRepo
}

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Create(category domain.Category) (*domain.Category, error) {
	query := `INSERT INTO categories (
			name,
			slug,
			description
		) VALUES(
			$1,
			$2,
			$3
		) RETURNING id;
	`

	row := r.db.QueryRow(query, category.Name, category.Slug, category.Description)
	err := row.Scan(&category.ID)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepo) List(page, limit int64) ([]*domain.Category, error) {

}

func (r *categoryRepo) Get(id int) (*domain.Category, error) {

}

func (r *categoryRepo) Update(category domain.Category) (*domain.Category, error) {

}

func (r *categoryRepo) Delete(id int) error {

}
