package main

import (
	"fmt"
	"github.com/erikrios/open-music-api-go-language/src/api/home"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	home.HomeRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		fmt.Println("Failed to start the server:", err)
	} else {
		fmt.Println("Server starting on port", 3000)
	}
}
