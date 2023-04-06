package controllers

import (
	"douglasdenny45.github.com/go/internal/infrastructure/controllers"
	"douglasdenny45.github.com/go/internal/infrastructure/factories/domain/services"
)

func MountGetUserController(controller controllers.Controller) controllers.ControllerInterface {
	return controllers.NewGetUserController(controller, services.MountGetUser())
}
