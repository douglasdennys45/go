package controllers

import (
	"github.com/douglasdennys45/go/internal/infrastructure/controllers"
	"github.com/douglasdennys45/go/internal/infrastructure/factories/infrastructure/database/postgres"
)

func MountDBTransactionController(controller controllers.ControllerInterface) controllers.ControllerInterface {
	return controllers.NewDBTransactionController(controller, postgres.MountPostgreSQLTransaction())
}
