package controllers

import (
	"douglasdenny45.github.com/go/internal/domain/services"
	"github.com/gofiber/fiber/v2"
)

type userDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddUserController struct {
	addUser services.AddUserInterface
}

func NewAddUserController(addUser services.AddUserInterface) Controller {
	return &AddUserController{addUser: addUser}
}

func (controller *AddUserController) Handle(ctx *fiber.Ctx) error {
	user := new(userDTO)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if err := controller.addUser.Execute(user.Name, user.Email, user.Password); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}
