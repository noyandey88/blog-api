package category

import "github.com/noyandey88/blog-api/domain"

type service struct {
	categoryRepo CategoryRepo
}

func NewService(categoryRepo CategoryRepo) Service {
	return &service{
		categoryRepo: categoryRepo,
	}
}

func (svc *service) Create(category domain.Category) (*domain.Category, error) {
	ctgry, err := svc.categoryRepo.Create(category)

	if err != nil {
		return nil, err
	}

	if ctgry == nil {
		return nil, nil
	}

	return ctgry, nil

}

func (svc *service) List() ([]*domain.Category, error) {
	categories, err := svc.categoryRepo.List()

	if err != nil {
		return nil, err
	}

	if categories == nil {
		return nil, nil
	}

	return categories, nil
}

func (svc *service) Get(id int) (*domain.Category, error) {
	category, err := svc.categoryRepo.Get(id)

	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, nil
	}

	return category, nil
}

func (svc *service) Update(category domain.Category) (*domain.Category, error) {
	ctgry, err := svc.categoryRepo.Update(category)

	if err != nil {
		return nil, err
	}

	if ctgry == nil {
		return nil, nil
	}

	return ctgry, nil
}

func (svc *service) Delete(id int) error {
	err := svc.categoryRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
