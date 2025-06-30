package organizations

import "github.com/adilsonchacon/goeli/entities"

type OrganizationData struct {
	Organization Organization `json:"data"`
}

type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Organizations struct {
	Data       []Organization      `json:"data"`
	Pagination entities.Pagination `json:"pagination"`
}
