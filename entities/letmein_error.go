package entities

type LetmeinError struct {
	Errors Detail `json:"errors"`
}

type Detail struct {
	Detail string `json:"detail"`
}

type LetmeinValidationError struct {
	Field   string
	Message []string
}
