package repositories

import (
	repository "douglasdenny45.github.com/go/internal/domain/repositories"
	"douglasdenny45.github.com/go/internal/infrastructure/database/postgres/repositories"
)

func MountUserRepository() repository.UserRepo {
	return repositories.NewUserRepository()
}
