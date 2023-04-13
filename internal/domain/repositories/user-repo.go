package repositories

import "github.com/douglasdennys45/go/internal/domain/entities"

type UserRepo interface {
	Create(*entities.User) error
	GetByID(id string) (*entities.User, error)
}
