package entities

type LetmeinError struct {
	Errors Detail `json:"errors"`
}

type Detail struct {
	Detail string `json:"detail"`
}
