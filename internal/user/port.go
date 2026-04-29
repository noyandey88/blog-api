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
	FindByEmail(email string) (*domain.User, error)
}
