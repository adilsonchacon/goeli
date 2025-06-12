package entities

type LetmeinMessage struct {
	Data Message `json:"data"`
}

type Message struct {
	Message string `json:"message"`
}
