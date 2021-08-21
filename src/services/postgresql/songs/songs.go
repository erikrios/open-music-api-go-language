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
	db, err := database.Db()
	if err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError("Something went wrong.")
	}

	statement := `SELECT id, title, year, performer, genre, duration, inserted_at, updated_at
					FROM songs`

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

// GetSong is a function to get a song in the database by id
func GetSong(id string) (*Song, errors.Error) {
	db, err := database.Db()
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		return nil, errors.NewInternalServerError("Something went wrong.")
	}
}

// UpdateSong is a function to update a song in the database by id
func UpdateSong(id string, year uint16, performer string, genre *string, duration *uint16) errors.Error {
	db, err := database.Db()
	if err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError("Something went wrong.")
	}

	statement := `UPDATE songs SET year = $2, performer = $3, genre = $4, duration = $5 WHERE id = $1`
	result, err := db.Exec(statement, id, year, performer, genre, duration)
	if err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError("Something went wrong.")
	}

	if count, err := result.RowsAffected(); err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
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
		fmt.Println(err)
		return false, errors.NewInternalServerError("Something went wrong.")
	}
}
