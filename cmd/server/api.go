package main

import (
	"fmt"
	"time"

	"douglasdenny45.github.com/go/internal/infrastructure/adapters"
	ctrl "douglasdenny45.github.com/go/internal/infrastructure/controllers"
	"douglasdenny45.github.com/go/internal/infrastructure/database/mysql"
	"douglasdenny45.github.com/go/internal/infrastructure/factories/infrastructure/controllers"
	"douglasdenny45.github.com/go/internal/infrastructure/middlewares"
	"douglasdenny45.github.com/go/pkg/env"
	"douglasdenny45.github.com/go/pkg/logger"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func init() {
	env, err := env.NewEnv(".")
	if err != nil {
		panic("env is nil")
	}
	_, err = mysql.NewMysqlConnect(env.DBDriver, env.DBDSN, env.DBPool)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	log := logger.NewLogger()
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
	log.Info("Server running on port 3000")
	app.Listen(":3000")
}
