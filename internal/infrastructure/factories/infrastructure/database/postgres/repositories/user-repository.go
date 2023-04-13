package repositories

import (
	repository "github.com/douglasdennys45/go/internal/domain/repositories"
	"github.com/douglasdennys45/go/internal/infrastructure/database/postgres/repositories"
)

func MountUserRepository() repository.UserRepo {
	return repositories.NewUserRepository()
}
