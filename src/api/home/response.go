package home

type Success struct {
	Status  string  `json:"status"`
	Message *string `json:"message"`
	Data    Data    `json:"data"`
}

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    *Data  `json:"data"`
}

type Data struct {
	Message string `json:"message"`
}
