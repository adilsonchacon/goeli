package apps

import "github.com/adilsonchacon/goeli/entities"

type DataApp struct {
	App App `json:"data"`
}

type App struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Apps struct {
	Apps       []App               `json:"data"`
	Pagination entities.Pagination `json:"pagination"`
}

type AppUsers struct {
	Users []AppUser `json:"data"`
}

type AppUser struct {
	ID   string `json:"id"`
	User User   `json:"user"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AppToken struct {
	ID        string  `json:"id"`
	AppID     string  `json:"app_id"`
	Token     *string `json:"token"`
	RevokedAt *string `json:"revoked_at"`
	RevokedBy *string `json:"revoked_by"`
	CreatedAt string  `json:"created_at"`
}

type AppTokens struct {
	AppTokens  []AppToken          `json:"data"`
	Pagination entities.Pagination `json:"pagination"`
}
