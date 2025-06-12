package entities

type LetmeinToken struct {
	Data Token `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}
