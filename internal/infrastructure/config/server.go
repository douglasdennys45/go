package config

import (
	"github.com/douglasdennys45/go/internal/infrastructure/adapters"
	ctrl "github.com/douglasdennys45/go/internal/infrastructure/controllers"
	"github.com/douglasdennys45/go/internal/infrastructure/factories/infrastructure/controllers"
	"github.com/douglasdennys45/go/internal/infrastructure/middlewares"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"time"
)

func NewServer() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message":   err.Error(),
				"status":    fiber.StatusInternalServerError,
				"error":     "Internal Server Error",
				"path":      ctx.Path(),
				"timestamp": time.Now(),
				"requestId": ctx.Locals("requestId"),
			})
		},
	})
	app.Use(middlewares.SetupRecover())
	app.Use(middlewares.SetupCors())
	app.Use(middlewares.SetupCompress())
	app.Use(middlewares.SetupLogger())
	app.Use(middlewares.SetupRequestId())
	app.Use(middlewares.SetupIdempotency())

	users := app.Group("/users")
	var controller ctrl.Controller
	users.Post("", adapters.RouteAdapter(controllers.MountAddUserController(controller)))
	users.Get(":id", adapters.RouteAdapter(controllers.MountGetUserController(controller)))

	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":   "route not found",
			"status":    fiber.StatusNotFound,
			"error":     "Not Found",
			"path":      ctx.Path(),
			"timestamp": time.Now(),
			"requestId": ctx.Locals("requestId"),
		})
	})
	return app
}
