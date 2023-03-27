package models

type Message struct {
	Text string `json:"message"`
}

type Response struct {
	Data    map[string]interface{} `json:"data"`
	Error   string                 `json:"error"`
	Message string                 `json:"message"`
}
