package organizations

type AdminUserDao interface {
	ListAdminUsers(orgID string, page, perPage int) (*AdminUsers, error)
	AddAdminUser(orgID, email string) (*AdminUser, error)
	RemoveAdminUser(orgID string, adminUserID string) error
}
