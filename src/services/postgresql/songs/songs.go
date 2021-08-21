package songs

import (
	"database/sql"
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

// AddSong is a function to add a song into the database
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

// GetSongs is a function to get songs in the database
func GetSongs() ([]Song, errors.Error) {
	statement := `SELECT id, title, year, performer, genre, duration, inserted_at, updated_at
					FROM songs`

	db, err := database.Db()
	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
		return nil, errors.NewInternalServerError("Something went wrong.")
	}

	rows, err := db.Query(statement)
	if err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError("Something went wrong.")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)

	var songs = make([]Song, 0)
	for rows.Next() {
		var song Song
		err := rows.Scan(&song.Id, &song.Title, &song.Year, &song.Performer, &song.Genre, &song.Duration, &song.InsertedAt, &song.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return nil, errors.NewInternalServerError("Something went wrong.")
		}
		songs = append(songs, song)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError("Something went wrong.")
	}

	return songs, nil
}
