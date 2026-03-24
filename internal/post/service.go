package post

import "github.com/noyandey88/blog-api/domain"

type service struct {
	postRepo PostRepo
}

func NewService(postRepo PostRepo) Service {
	return &service{
		postRepo: postRepo,
	}
}

func (svc *service) Create(post domain.Post) (*domain.Post, error) {
	return svc.postRepo.Create(post)
}

func (svc *service) List(page, limit int64) ([]*domain.Post, error) {
	return svc.postRepo.List(page, limit)
}

func (svc *service) Count() (int64, error) {
	return svc.postRepo.Count()
}

func (svc *service) Get(id int) (*domain.Post, error) {
	return svc.postRepo.Get(id)
}

func (svc *service) Update(post domain.Post) (*domain.Post, error) {
	return svc.postRepo.Update(post)
}

func (svc *service) Delete(id int) error {
	return svc.postRepo.Delete(id)
}
