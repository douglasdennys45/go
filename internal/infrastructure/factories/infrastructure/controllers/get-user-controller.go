package controllers

import (
	"douglasdenny45.github.com/go/internal/infrastructure/controllers"
	"douglasdenny45.github.com/go/internal/infrastructure/factories/domain/services"
)

func MountGetUserController() controllers.Controller {
	return controllers.NewGetUserController(services.MountGetUser())
}
