package songs

import (
	"database/sql"
	"github.com/erikrios/open-music-api-go-language/src/api/songs/payloads"
	"github.com/erikrios/open-music-api-go-language/src/api/songs/response"
	"github.com/erikrios/open-music-api-go-language/src/errors"
	service "github.com/erikrios/open-music-api-go-language/src/services/postgresql/songs"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/validator.v2"
)

func postSongs(c *fiber.Ctx) error {
	payload := new(payloads.Payload)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid payload or request body.",
		})
	}

	if err := validator.Validate(*payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	title, year, performer := payload.Title, payload.Year, payload.Performer
	genre, duration := sql.NullString{}, sql.NullInt16{}

	if payload.Genre == nil {
		genre.Valid = false
	} else {
		genre.Valid = true
		genre.String = *payload.Genre
	}
	if payload.Duration == nil {
		duration.Valid = false
	} else {
		duration.Valid = true
		duration.Int16 = int16(*payload.Duration)
	}

	id, err := service.AddSong(title, year, performer, genre, duration)
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
	results := make([]response.Simple, 0)

	allSongs, err := service.GetSongs()
	if err != nil {
		return errors.ErrorHandler(err, c)
	}

	for _, song := range allSongs {
		result := response.Simple{
			Id:        song.Id,
			Title:     song.Title,
			Performer: song.Performer,
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

	var genre = new(string)
	var duration = new(uint16)

	if song.Genre.Valid {
		*genre = song.Genre.String
	} else {
		genre = nil
	}

	if song.Duration.Valid {
		*duration = uint16(song.Duration.Int16)
	} else {
		duration = nil
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"song": response.Full{
				Id:         song.Id,
				Title:      song.Title,
				Year:       song.Year,
				Performer:  song.Performer,
				Genre:      genre,
				Duration:   duration,
				InsertedAt: song.InsertedAt,
				UpdatedAt:  song.UpdatedAt,
			},
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

	if err := validator.Validate(*payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	title, year, performer := payload.Title, payload.Year, payload.Performer
	genre, duration := sql.NullString{}, sql.NullInt16{}

	if payload.Genre == nil {
		genre.Valid = false
	} else {
		genre.Valid = true
		genre.String = *payload.Genre
	}
	if payload.Duration == nil {
		duration.Valid = false
	} else {
		duration.Valid = true
		duration.Int16 = int16(*payload.Duration)
	}

	if err := service.UpdateSong(id, title, year, performer, genre, duration); err != nil {
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
