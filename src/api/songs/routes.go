package songs

import "github.com/gofiber/fiber/v2"

func Routes(a *fiber.App) {
	router := a.Group("/")
	router.Post("/songs", postSongs)
	router.Get("/songs", getSongs)
	router.Get("/songs/:id", getSong)
	router.Put("/songs/:id", putSong)
	router.Delete("/songs/:id", deleteSong)
}
