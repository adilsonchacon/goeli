package organizations

import "github.com/adilsonchacon/goeli/entities"

type AdminUsers struct {
	Data       []AdminUser         `json:"data"`
	Pagination entities.Pagination `json:"pagination"`
}

type AdminUser struct {
	ID   string `json:"id"`
	User User   `json:"user"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
