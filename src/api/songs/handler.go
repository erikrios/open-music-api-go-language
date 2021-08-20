package songs

import (
	"github.com/erikrios/open-music-api-go-language/src/api/songs/payloads"
	service "github.com/erikrios/open-music-api-go-language/src/services/inmemory/songs"
	"github.com/erikrios/open-music-api-go-language/src/validation/songs"
	"github.com/gofiber/fiber/v2"
)

func postSongs(c *fiber.Ctx) error {
	payload := new(payloads.Payload)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if errors := songs.Validate(*payload); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": fiber.Map{
				"errors": errors,
			},
		})
	}

	id := service.AddSong(payload.Title, payload.Year, payload.Performer, payload.Genre, payload.Duration)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Lagu berhasil ditambahkan",
		"data": fiber.Map{
			"songId": id,
		},
	})
}

func getSongs(c *fiber.Ctx) error {
	results := make([]fiber.Map, 0)

	for _, song := range service.GetSongs() {
		result := fiber.Map{
			"id":        song.Id,
			"title":     song.Title,
			"performer": song.Performer,
		}
		results = append(results, result)
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"songs": results,
		},
	})
}

func getSong(c *fiber.Ctx) error {
	id := c.Params("id")
	song, err := service.GetSong(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Message(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"song": song,
		},
	})
}
