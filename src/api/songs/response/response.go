package response

type Simple struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Performer string `json:"performer"`
}

type Full struct {
	Id         string  `json:"id"`
	Title      string  `json:"title"`
	Year       uint16  `json:"year"`
	Performer  string  `json:"performer"`
	Genre      *string `json:"genre"`
	Duration   *uint16 `json:"duration"`
	InsertedAt string  `json:"insertedAt"`
	UpdatedAt  string  `json:"updatedAt"`
}
