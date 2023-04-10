package services

import (
	"douglasdenny45.github.com/go/internal/domain/services"
	"douglasdenny45.github.com/go/internal/infrastructure/factories/infrastructure/database/postgres/repositories"
)

func MountGetUser() services.GetUserInterface {
	return services.NewGetUser(repositories.MountUserRepository())
}
