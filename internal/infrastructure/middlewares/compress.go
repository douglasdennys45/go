package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func SetupCompress() fiber.Handler {
	return compress.New(compress.Config{Level: compress.LevelBestSpeed})
}
