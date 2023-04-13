package services

import (
	"github.com/douglasdennys45/go/internal/domain/entities"
	"github.com/douglasdennys45/go/internal/domain/repositories"
)

type GetUserInterface interface {
	Execute(id string) (*entities.User, error)
}

type GetUser struct {
	userRepo repositories.UserRepo
}

func NewGetUser(userRepo repositories.UserRepo) GetUserInterface {
	return &GetUser{userRepo: userRepo}
}

func (a *GetUser) Execute(id string) (*entities.User, error) {
	user, err := a.userRepo.GetByID(id)
	return user, err
}
