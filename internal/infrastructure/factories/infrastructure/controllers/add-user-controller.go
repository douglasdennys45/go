package controllers

import (
	"github.com/douglasdennys45/go/internal/infrastructure/controllers"
	"github.com/douglasdennys45/go/internal/infrastructure/factories/domain/services"
)

func MountAddUserController(controller controllers.Controller) controllers.ControllerInterface {
	return MountDBTransactionController(controllers.NewAddUserController(controller, services.MountAddUser()))
}
