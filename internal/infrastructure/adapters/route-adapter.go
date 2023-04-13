package adapters

import (
	"github.com/douglasdennys45/go/internal/infrastructure/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteAdapter(controller controllers.ControllerInterface) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctrl := controllers.Controller{Ctx: ctx}
		response := ctrl.Handle(controller)
		return ctx.Status(response.Status).JSON(fiber.Map{
			"data":      response.Data,
			"message":   response.Message,
			"status":    response.Status,
			"path":      response.Path,
			"timestamp": response.Timestamp,
			"requestId": ctx.Locals("requestId"),
		})
	}
}
