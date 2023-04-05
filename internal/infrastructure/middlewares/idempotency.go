package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
)

func SetupIdempotency() fiber.Handler {
	return idempotency.New()
}
