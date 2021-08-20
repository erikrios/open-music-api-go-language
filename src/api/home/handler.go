package home

import "github.com/gofiber/fiber/v2"

func getHome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": nil,
		"data": fiber.Map{
			"message": "Hello, World!",
		},
	})
}
