package entities

type Pagination struct {
	Count   int      `json:"count"`
	First   int      `json:"first"`
	Last    int      `json:"last"`
	Next    *int     `json:"next"`
	Prev    *int     `json:"prev"`
	Page    int      `json:"page"`
	PerPage int      `json:"per_page"`
	Serie   []string `json:"serie"`
}
