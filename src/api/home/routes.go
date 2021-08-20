package home

import "github.com/gofiber/fiber/v2"

func HomeRoutes(a *fiber.App) {
	router := a.Group("/")
	router.Get("/", getHome)
}
