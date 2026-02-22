package tag

import "github.com/noyandey88/blog-api/domain"

type service struct {
	tagRepo TagRepo
}

func NewService(tagRepo TagRepo) Service {
	return &service{
		tagRepo: tagRepo,
	}
}

func (s *service) Create(tag *domain.Tag) (*domain.Tag, error) {
	tag, err := s.tagRepo.Create(tag)
	if err != nil {
		return nil, err
	}

	if tag == nil {
		return nil, nil
	}

	return tag, nil
}

func (s *service) List() ([]*domain.Tag, error) {
	tags, err := s.tagRepo.List()
	if err != nil {
		return nil, err
	}

	if tags == nil {
		return nil, nil
	}

	return tags, nil
}

func (s *service) Get(id int) (*domain.Tag, error) {
	tag, err := s.tagRepo.Get(id)
	if err != nil {
		return nil, err
	}

	if tag == nil {
		return nil, nil
	}

	return tag, nil
}

func (s *service) Update(id int, tag *domain.Tag) (*domain.Tag, error) {
	tag, err := s.tagRepo.Update(id, tag)
	if err != nil {
		return nil, err
	}

	if tag == nil {
		return nil, nil
	}

	return tag, nil
}

func (s *service) Delete(id int) error {
	err := s.tagRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
