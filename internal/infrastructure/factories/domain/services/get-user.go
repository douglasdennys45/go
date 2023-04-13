package services

import (
	"github.com/douglasdennys45/go/internal/domain/services"
	"github.com/douglasdennys45/go/internal/infrastructure/factories/infrastructure/database/postgres/repositories"
)

func MountGetUser() services.GetUserInterface {
	return services.NewGetUser(repositories.MountUserRepository())
}
