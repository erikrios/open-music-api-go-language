package songs

import (
	"fmt"
	"github.com/aidarkhanov/nanoid"
	"github.com/erikrios/open-music-api-go-language/src/errors"
	"time"
)

type Song struct {
	Id         string        `json:"id"`
	Title      string        `json:"title"`
	Year       uint16        `json:"year"`
	Performer  string        `json:"performer"`
	Genre      string        `json:"genre"`
	Duration   time.Duration `json:"duration"`
	InsertedAt string        `json:"insertedAt"`
	UpdatedAt  string        `json:"updatedAt"`
}

type Payload struct {
	Title     string        `json:"title"`
	Year      uint16        `json:"year"`
	Performer string        `json:"performer"`
	Genre     string        `json:"genre"`
	Duration  time.Duration `json:"duration"`
}

var songs = make([]Song, 0)

func AddSong(payload Payload) string {
	id, _ := nanoid.Generate(nanoid.DefaultAlphabet, 16)
	insertedAt := time.Now().Format(time.RFC3339)
	updatedAt := insertedAt

	song := Song{
		Id:         id,
		Title:      payload.Title,
		Year:       payload.Year,
		Performer:  payload.Performer,
		Genre:      payload.Genre,
		Duration:   payload.Duration,
		InsertedAt: insertedAt,
		UpdatedAt:  updatedAt,
	}
	songs = append(songs, song)
	return id
}

func GetSongs() []map[string]interface{} {
	results := make([]map[string]interface{}, 0)

	for _, song := range songs {
		result := map[string]interface{}{
			"id":        song.Id,
			"title":     song.Title,
			"performer": song.Performer,
		}

		results = append(results, result)
	}

	return results
}

func GetSong(id string) (*Song, errors.Error) {
	for _, song := range songs {
		if id == song.Id {
			return &song, nil
		}
	}
	return nil, errors.NewNotFound(fmt.Sprintf("Song with id %s not found.", id))
}

func UpdateSong(id string, payload Payload) errors.Error {
	song, err := GetSong(id)

	if err != nil {
		return err
	}

	song.Title = payload.Title
	song.Year = payload.Year
	song.Performer = payload.Performer
	song.Genre = payload.Genre
	song.Duration = payload.Duration
	return nil
}

func DeleteSong(id string, payload Payload) errors.Error {
	for i, song := range songs {
		if id == song.Id {
			songs = remove(songs, i)
			return nil
		}
	}

	return errors.NewNotFound(fmt.Sprintf("Song with id %s not found.", id))
}

func remove(s []Song, i int) []Song {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
