package middleware

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(a *fiber.App) {
	swaggerMiddleware(a)
	corsMiddleware(a)
	compressMiddleware(a)
	eTagMiddleware(a)
	loggerMiddleware(a)
}

func swaggerMiddleware(a *fiber.App) {
	a.Get("/swagger/*", swagger.Handler)
}

func corsMiddleware(a *fiber.App) {
	a.Use(cors.New())
}

func compressMiddleware(a *fiber.App) {
	a.Use(compress.New())
}

func eTagMiddleware(a *fiber.App) {
	a.Use(etag.New())
}

func loggerMiddleware(a *fiber.App) {
	a.Use(logger.New())
}
