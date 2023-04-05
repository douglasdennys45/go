package controllers

import (
	"douglasdenny45.github.com/go/internal/infrastructure/controllers"
	"douglasdenny45.github.com/go/internal/infrastructure/factories/domain/services"
)

func MountAddUserController() controllers.Controller {
	return controllers.NewAddUserController(services.MountAddUser())
}
