package songs

import (
	"github.com/erikrios/open-music-api-go-language/src/api/songs/payloads"
	"github.com/erikrios/open-music-api-go-language/src/errors"
	service "github.com/erikrios/open-music-api-go-language/src/services/postgresql/songs"
	"github.com/erikrios/open-music-api-go-language/src/validation/songs"
	"github.com/gofiber/fiber/v2"
)

func postSongs(c *fiber.Ctx) error {
	payload := new(payloads.Payload)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid payload or request body.",
		})
	}

	if err := songs.Validate(*payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": fiber.Map{
				"err": err,
			},
		})
	}

	id, err := service.AddSong(payload.Title, payload.Year, payload.Performer, payload.Genre, payload.Duration)
	if err != nil {
		return errors.ErrorHandler(err, c)
	}

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

	allSongs, err := service.GetSongs()
	if err != nil {
		return errors.ErrorHandler(err, c)
	}

	for _, song := range allSongs {
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
		return errors.ErrorHandler(err, c)
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"song": song,
		},
	})
}

func putSong(c *fiber.Ctx) error {
	id := c.Params("id")
	payload := new(payloads.Payload)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid payload or request body.",
		})
	}

	if err := songs.Validate(*payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": fiber.Map{
				"err": err,
			},
		})
	}

	if err := service.UpdateSong(id, payload.Title, payload.Year, payload.Performer, payload.Genre, payload.Duration); err != nil {
		return errors.ErrorHandler(err, c)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "lagu berhasil diperbarui",
	})
}

func deleteSong(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := service.DeleteSong(id); err != nil {
		return errors.ErrorHandler(err, c)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "lagu berhasil dihapus",
	})
}
