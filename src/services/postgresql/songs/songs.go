package songs

import (
	"fmt"
	"github.com/aidarkhanov/nanoid"
	"github.com/erikrios/open-music-api-go-language/src/database"
	"github.com/erikrios/open-music-api-go-language/src/errors"
	"time"
)

type Song struct {
	Id         string  `json:"id"`
	Title      string  `json:"title"`
	Year       uint16  `json:"year"`
	Performer  string  `json:"performer"`
	Genre      *string `json:"genre"`
	Duration   *uint16 `json:"duration"`
	InsertedAt string  `json:"insertedAt"`
	UpdatedAt  string  `json:"updatedAt"`
}

func AddSong(title string, year uint16, performer string, genre *string, duration *uint16) (string, errors.Error) {
	nanoidId, _ := nanoid.Generate(nanoid.DefaultAlphabet, 16)
	id := fmt.Sprintf("song-%s", nanoidId)
	insertedAt := time.Now().Format(time.RFC3339)

	statement := `INSERT INTO songs (id, title, year, performer, genre, duration, inserted_at, updated_at)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $7)
					RETURNING id`

	db, err := database.Db()
	if err != nil {
		fmt.Println(err)
		return "", errors.NewInternalServerError("Something went wrong.")
	}

	row := db.QueryRow(statement, id, title, year, performer, genre, duration, insertedAt)
	if err := row.Scan(&id); err != nil {
		fmt.Println(err)
		return "", errors.NewInternalServerError("Something went wrong.")
	}
	return id, nil
}
