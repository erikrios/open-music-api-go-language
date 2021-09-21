package songs

import "github.com/gofiber/fiber/v2"

func Routes(a *fiber.App) {
	router := a.Group("/songs")
	router.Post("/", postSongs)
	router.Get("/", getSongs)
	router.Get("/:id", getSong)
	router.Put("/:id", putSong)
	router.Delete("/:id", deleteSong)
}
