package repositories

import (
	repository "douglasdenny45.github.com/go/internal/domain/repositories"
	"douglasdenny45.github.com/go/internal/infrastructure/database/mysql/repositories"
	"douglasdenny45.github.com/go/internal/infrastructure/factories/infrastructure/database/mysql"
)

func MountUserRepository() repository.UserRepo {
	return repositories.NewUserRepository(mysql.MountMySQLConnect())
}
