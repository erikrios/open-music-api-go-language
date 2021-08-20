package home

import (
	_ "github.com/erikrios/open-music-api-go-language/docs"
	"github.com/gofiber/fiber/v2"
)

// GetHome godoc
// @Summary Get hello world message
// @Description get hello world message from the server
// @Accept  json
// @Produce  json
// @Success 200 {object} Success
// @Failure 404,500 {object} Error
// @Failure default {object} Error
// @Router / [get]
func getHome(c *fiber.Ctx) error {
	return c.JSON(Success{
		Status:  "success",
		Message: nil,
		Data: Data{
			Message: "Hello, World!",
		},
	})
}
