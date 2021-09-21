package payloads

type Payload struct {
	Title     string  `json:"title" validate:"nonzero,min=1,max=100"`
	Year      uint16  `json:"year" validate:"nonzero,min=1970,max=2025"`
	Performer string  `json:"performer" validate:"nonzero,min=1,max=50"`
	Genre     *string `json:"genre" validate:"min=1,max=50"`
	Duration  *uint16 `json:"duration" validate:"min=1,max=20000"`
}
