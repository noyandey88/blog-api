package user

import (
	"errors"

	"github.com/noyandey88/blog-api/domain"
	"github.com/noyandey88/blog-api/utils"
)

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	usr, err := svc.userRepo.Create(user)

	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil
}

func (svc *service) Find(email string, password string) (*domain.User, error) {
	usr, err := svc.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	err = utils.ComparePassword(usr.Password, password)

	if err != nil {
		return nil, errors.New("Invalid credentials")
	}

	return usr, nil
}
