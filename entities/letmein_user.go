package entities

type DataUser struct {
	User User `json:"data"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
	Language string `json:"language"`
	Timezone string `json:"timezone"`
}
