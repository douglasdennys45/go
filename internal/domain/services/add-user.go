package services

import (
	"douglasdenny45.github.com/go/internal/domain/entities"
	"douglasdenny45.github.com/go/internal/domain/repositories"
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
