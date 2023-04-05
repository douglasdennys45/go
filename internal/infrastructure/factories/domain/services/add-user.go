package services

import (
	"douglasdenny45.github.com/go/internal/domain/services"
	"douglasdenny45.github.com/go/internal/infrastructure/factories/infrastructure/database/mysql/repositories"
)

func MountAddUser() services.AddUserInterface {
	return services.NewAddUser(repositories.MountUserRepository())
}
