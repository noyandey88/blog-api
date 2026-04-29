package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/noyandey88/blog-api/domain"
	"github.com/noyandey88/blog-api/internal/tag"
)

type TagRepo interface {
	tag.TagRepo
}

type tagRepo struct {
	db *sqlx.DB
}

func NewTagRepo(db *sqlx.DB) TagRepo {
	return &tagRepo{
		db: db,
	}
}

func (r *tagRepo) Create(tag domain.Tag) (*domain.Tag, error) {
	query := `INSERT INTO tags (
			name,
			slug
		) VALUES(
			$1,
			$2
		) RETURNING id, created_at, updated_at;`

	row := r.db.QueryRow(query, tag.Name, tag.Slug)
	err := row.Scan(
		&tag.ID,
		&tag.CreatedAt,
		&tag.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r *tagRepo) List() ([]*domain.Tag, error) {
	query := `
		SELECT
			id,
			name,
			slug,
			created_at,
			updated_at
		FROM tags;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*domain.Tag

	for rows.Next() {
		var tag domain.Tag
		err := rows.Scan(&tag.ID, &tag.Name, &tag.Slug, &tag.CreatedAt, &tag.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}

	return tags, nil
}

func (r *tagRepo) Get(id int) (*domain.Tag, error) {
	query := `SELECT * FROM tags WHERE id = $1;`

	row := r.db.QueryRow(query, id)
	var tag domain.Tag
	err := row.Scan(&tag.ID, &tag.Name, &tag.Slug, &tag.CreatedAt, &tag.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r *tagRepo) Update(tag domain.Tag) (*domain.Tag, error) {
	query := `UPDATE tags SET
			name = $1,
			updated_at = $2
			WHERE id = $3
			RETURNING id,
			name,
			slug,
			created_at,
			updated_at;`

	row := r.db.QueryRow(query, tag.Name, tag.UpdatedAt, tag.ID)
	err := row.Scan(&tag.ID, &tag.Name, &tag.Slug, &tag.CreatedAt, &tag.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r *tagRepo) Delete(id int) error {
	query := `DELETE FROM tags WHERE id = $1;`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
