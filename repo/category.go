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

func (r *categoryRepo) List() ([]*domain.Category, error) {
	query := `SELECT id, name, slug, description FROM categories;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*domain.Category

	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Slug, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (r *categoryRepo) Get(id int) (*domain.Category, error) {
	query := `SELECT id, name, slug, description FROM categories WHERE id = $1;`

	row := r.db.QueryRow(query, id)
	var category domain.Category
	err := row.Scan(&category.ID, &category.Name, &category.Slug, &category.Description)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepo) Update(category domain.Category) (*domain.Category, error) {
	query := `UPDATE categories SET name = $1, slug = $2, description = $3 WHERE id = $4 RETURNING id;`

	row := r.db.QueryRow(query, category.Name, category.Slug, category.Description, category.ID)
	err := row.Scan(&category.ID)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepo) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1;`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
