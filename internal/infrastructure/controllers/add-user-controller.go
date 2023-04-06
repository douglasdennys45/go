package controllers

import (
	"time"

	"douglasdenny45.github.com/go/internal/domain/services"
	"github.com/gofiber/fiber/v2"
)

type userDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddUserController struct {
	Controller
	services.AddUserInterface
}

func NewAddUserController(controller Controller, addUser services.AddUserInterface) ControllerInterface {
	return &AddUserController{controller, addUser}
}

func (controller *AddUserController) perform(ctx *fiber.Ctx) Response {
	user := new(userDTO)
	if err := ctx.BodyParser(user); err != nil {
		return Response{
			Message:   err.Error(),
			Status:    fiber.StatusInternalServerError,
			Path:      "POST /users",
			Timestamp: time.Now(),
		}
	}
	if err := controller.Execute(user.Name, user.Email, user.Password); err != nil {
		return Response{
			Message:   err.Error(),
			Status:    fiber.StatusBadRequest,
			Path:      "POST /users",
			Timestamp: time.Now(),
		}
	}
	return Response{
		Message:   "User created successfully",
		Status:    fiber.StatusCreated,
		Path:      "POST /users",
		Timestamp: time.Now(),
	}
}
