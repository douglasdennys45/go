package controllers

import (
	"douglasdenny45.github.com/go/internal/infrastructure/controllers"
	"douglasdenny45.github.com/go/internal/infrastructure/factories/domain/services"
)

func MountAddUserController(controller controllers.Controller) controllers.ControllerInterface {
	return controllers.NewAddUserController(controller, services.MountAddUser())
}
