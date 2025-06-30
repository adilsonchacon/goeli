package organizations

type OrganizationDao interface {
	List(page, perPage int) (*Organizations, error)
	Find(id string) (*Organization, error)
	Create(newOrganization Organization) (Organization, error)
	Update(organization Organization) error
	Delete(id string) error
	ListAdminUsers(id string, page, perPage int) (OrganizationAdminUsers, error)
	AddAdminUsers(user User) (*OrganizationAdminUser, error)
	RemoveAdminUsers(id string, adminUserID string) error
}
