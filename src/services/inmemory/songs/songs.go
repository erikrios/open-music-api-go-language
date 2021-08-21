package songs

import (
	"fmt"
	"github.com/aidarkhanov/nanoid"
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

var songs = make([]Song, 0)

func AddSong(title string, year uint16, performer string, genre *string, duration *uint16) string {
	nanoidId, _ := nanoid.Generate(nanoid.DefaultAlphabet, 16)
	id := fmt.Sprintf("song-%s", nanoidId)
	insertedAt := time.Now().Format(time.RFC3339)
	updatedAt := insertedAt

	song := Song{
		Id:         id,
		Title:      title,
		Year:       year,
		Performer:  performer,
		Genre:      genre,
		Duration:   duration,
		InsertedAt: insertedAt,
		UpdatedAt:  updatedAt,
	}
	songs = append(songs, song)
	return id
}

func GetSongs() []Song {
	return songs
}

func GetSong(id string) (*Song, errors.Error) {
	for _, song := range songs {
		if id == song.Id {
			return &song, nil
		}
	}
	return nil, errors.NewNotFound(fmt.Sprintf("Song with id %s not found.", id))
}

func UpdateSong(id string, title string, year uint16, performer string, genre *string, duration *uint16) errors.Error {
	for i, song := range songs {
		if id == song.Id {
			songs[i].Title = title
			songs[i].Year = year
			songs[i].Performer = performer
			songs[i].Genre = genre
			songs[i].Duration = duration
			return nil
		}
	}
	return errors.NewNotFound(fmt.Sprintf("Song with id %s not found.", id))
}

func DeleteSong(id string) errors.Error {
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
