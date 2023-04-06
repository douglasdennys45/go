package controllers

import (
	"time"

	"douglasdenny45.github.com/go/internal/domain/services"
	"github.com/gofiber/fiber/v2"
)

type GetUserController struct {
	Controller
	services.GetUserInterface
}

func NewGetUserController(controller Controller, getUser services.GetUserInterface) ControllerInterface {
	return &GetUserController{controller, getUser}
}

func (controller *GetUserController) perform(ctx *fiber.Ctx) Response {
	id := ctx.Params("id")
	users, err := controller.Execute(id)
	if err != nil {
		return Response{
			Data:      nil,
			Message:   err.Error(),
			Status:    fiber.StatusNotFound,
			Path:      "GET /users/" + id,
			Timestamp: time.Now(),
		}
	}
	return Response{
		Data:      users,
		Message:   "User found successfully",
		Status:    fiber.StatusOK,
		Path:      "/users/" + id,
		Timestamp: time.Now(),
	}
}
