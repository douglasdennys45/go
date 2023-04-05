package repositories

import "douglasdenny45.github.com/go/internal/domain/entities"

type UserRepo interface {
	Create(*entities.User) error
	GetByID(id string) (*entities.User, error)
}
