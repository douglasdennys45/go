package controllers

import (
	"douglasdenny45.github.com/go/internal/domain/services"
	"github.com/gofiber/fiber/v2"
)

type GetUserController struct {
	getUser services.GetUserInterface
}

func NewGetUserController(getUser services.GetUserInterface) Controller {
	return &GetUserController{getUser: getUser}
}

func (controller *GetUserController) Handle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	users, err := controller.getUser.Execute(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(users)
}
