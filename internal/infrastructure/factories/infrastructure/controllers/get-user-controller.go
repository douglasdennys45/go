package controllers

import (
	"github.com/douglasdennys45/go/internal/infrastructure/controllers"
	"github.com/douglasdennys45/go/internal/infrastructure/factories/domain/services"
)

func MountGetUserController(controller controllers.Controller) controllers.ControllerInterface {
	return controllers.NewGetUserController(controller, services.MountGetUser())
}
