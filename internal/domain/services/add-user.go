package services

import (
	"github.com/douglasdennys45/go/internal/domain/entities"
	"github.com/douglasdennys45/go/internal/domain/repositories"
)

type AddUserInterface interface {
	Execute(name, email, password string) error
}

type AddUser struct {
	userRepo repositories.UserRepo
}

func NewAddUser(userRepo repositories.UserRepo) AddUserInterface {
	return &AddUser{userRepo: userRepo}
}

func (a *AddUser) Execute(name, email, password string) error {
	user, err := entities.NewUser(name, email, password)
	if err != nil {
		return err
	}
	return a.userRepo.Create(user)
}
