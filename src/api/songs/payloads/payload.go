package payloads

type Payload struct {
	Title     string  `json:"title" validate:"gte=1,lte=100,required"`
	Year      uint16  `json:"year" validate:"min=1970,max=2025,required"`
	Performer string  `json:"performer" validate:"gte=1,lte=50,required"`
	Genre     *string `json:"genre" validate:"omitempty,gte=1,lte=50"`
	Duration  *uint16 `json:"duration" validate:"omitempty,min=1,max=20000"`
}
