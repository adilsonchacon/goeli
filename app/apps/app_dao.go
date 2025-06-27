package apps

type AppDao interface {
	Create(newApp App) (App, error)
	List(organizationID string, page, perPage int) (Apps, error)
	Find(organizationID string, id string) (App, error)
	Update(app App) (App, error)
	Delete(organizationID string, id string) error
	Users(organizationID string, AppID string, page, perPage int) (AppUsers, error)
	AddUser(organizationID string, AppID string, user User) (AppUser, error)
	RemoveUser(organizationID string, AppID string, appUserID string) error
	CreateToken(organizationID string, appID string) (AppToken, error)
	ListTokens(organizationID string, appID string) (AppTokens, error)
	FindToken(organizationID string, appID, id string) (AppToken, error)
	RevokeToken(organizationID string, appID, id string) error
}
