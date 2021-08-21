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
		return "", errors.NewInternalServerError("Something went wrong.")
	}

	row := db.QueryRow(statement, id, title, year, performer, genre, duration, insertedAt)
	if err := row.Scan(&id); err != nil {
		return "", errors.NewInternalServerError("Something went wrong.")
	}
	return id, nil
}

// GetSongs is a function to get songs in the database
func GetSongs() ([]Song, errors.Error) {
	db, err := database.Db()
	if err != nil {
		return nil, errors.NewInternalServerError("Something went wrong.")
	}

	statement := `SELECT id, title, year, performer, genre, duration, inserted_at, updated_at
					FROM songs`

	rows, err := db.Query(statement)
	if err != nil {
		return nil, errors.NewInternalServerError("Something went wrong.")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	var songs = make([]Song, 0)
	for rows.Next() {
		var song Song
		err := rows.Scan(&song.Id, &song.Title, &song.Year, &song.Performer, &song.Genre, &song.Duration, &song.InsertedAt, &song.UpdatedAt)
		if err != nil {
			return nil, errors.NewInternalServerError("Something went wrong.")
		}
		songs = append(songs, song)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.NewInternalServerError("Something went wrong.")
	}

	return songs, nil
}

// GetSong is a function to get a song in the database by id
func GetSong(id string) (*Song, errors.Error) {
	db, err := database.Db()
	if err != nil {
		return nil, errors.NewInternalServerError("Something went wrong.")
	}

	statement := `SELECT id, title, year, performer, genre, duration, inserted_at, updated_at
					FROM songs
					WHERE id = $1`

	row := db.QueryRow(statement, id)

	var song Song
	switch row.Scan(
		&song.Id,
		&song.Title,
		&song.Year,
		&song.Performer,
		&song.Genre,
		&song.Duration,
		&song.InsertedAt,
		&song.UpdatedAt,
	) {
	case sql.ErrNoRows:
		return nil, errors.NewNotFound(fmt.Sprintf("Song with id %s not found.", id))
	case nil:
		return &song, nil
	default:
		return nil, errors.NewInternalServerError("Something went wrong.")
	}
}

// UpdateSong is a function to update a song in the database by id
func UpdateSong(id string, title string, year uint16, performer string, genre *string, duration *uint16) errors.Error {
	db, err := database.Db()
	if err != nil {
		return errors.NewInternalServerError("Something went wrong.")
	}

	statement := `UPDATE songs SET title = $2, year = $3, performer = $4, genre = $5, duration = $6 WHERE id = $1`
	result, err := db.Exec(statement, id, title, year, performer, genre, duration)
	if err != nil {
		return errors.NewInternalServerError("Something went wrong.")
	}

	if count, err := result.RowsAffected(); err != nil {
		return errors.NewInternalServerError("Something went wrong.")
	} else if count < 1 {
		return errors.NewNotFound(fmt.Sprintf("Song with id %s not found.", id))
	}

	return nil
}

// DeleteSong is a function to delete a song in the database by id
func DeleteSong(id string) errors.Error {
	db, err := database.Db()
	if err != nil {
		return errors.NewInternalServerError("Something went wrong.")
	}

	statement := `DELETE FROM songs WHERE id = $1`
	result, err := db.Exec(statement, id)
	if err != nil {
		return errors.NewInternalServerError("Something went wrong.")
	}

	if count, err := result.RowsAffected(); err != nil {
		return errors.NewInternalServerError("Something went wrong.")
	} else if count < 1 {
		return errors.NewNotFound(fmt.Sprintf("Song with id %s not found.", id))
	}

	return nil
}

// exists is a function to check the existence of the song in the database
func exists(id string) (bool, errors.Error) {
	db, err := database.Db()
	if err != nil {
		return false, errors.NewInternalServerError("Something went wrong.")
	}

	statement := `SELECT COUNT (id) FROM songs WHERE id = $1`
	row := db.QueryRow(statement, id)

	var count uint8
	switch err := row.Scan(&count); err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, errors.NewInternalServerError("Something went wrong.")
	}
}
