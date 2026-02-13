package user

import (
	"github.com/noyandey88/blog-api/domain"
	userHandler "github.com/noyandey88/blog-api/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email string, password string) (*domain.User, error)
}
