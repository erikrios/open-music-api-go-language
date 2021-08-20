package home

import "github.com/gofiber/fiber/v2"

func Routes(a *fiber.App) {
	router := a.Group("/")
	router.Get("/", getHome)
}
