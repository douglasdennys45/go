package controllers

import (
	"douglasdenny45.github.com/go/internal/infrastructure/controllers"
	"douglasdenny45.github.com/go/internal/infrastructure/factories/infrastructure/database/postgres"
)

func MountDBTransactionController(controller controllers.ControllerInterface) controllers.ControllerInterface {
	return controllers.NewDBTransactionController(controller, postgres.MountPostgreSQLTransaction())
}
