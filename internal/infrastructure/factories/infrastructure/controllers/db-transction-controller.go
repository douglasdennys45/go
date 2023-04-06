package controllers

import "douglasdenny45.github.com/go/internal/infrastructure/controllers"

func MountDBTransactionController(controller controllers.ControllerInterface) controllers.DBTransactionController {
	return controllers.DBTransactionController{
		controller,
		nil,
	}
}
