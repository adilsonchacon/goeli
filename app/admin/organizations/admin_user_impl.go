package organizations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adilsonchacon/goeli/lib/letmeinerr"
	"github.com/adilsonchacon/goeli/lib/restapi"
)

func (letmein *OrganizationRepo) ListAdminUsers(orgID string, page, perPage int) (*AdminUsers, error) {
	url := letmein.Repo.BaseURL + "/rest/admin/organizations/" + orgID + "/admin_users?page=" + strconv.Itoa(page) + "&perPage=" + strconv.Itoa(perPage)
	req := restapi.New(url, http.MethodGet)
	req.AddHeader("Authorization", fmt.Sprintf("Bearer %s", letmein.Repo.SessionToken))
	statusCode, body, err := req.DoRequest()

	if err != nil {
		return nil, fmt.Errorf("error requesting list of organization's admin users: %s", err)
	}

	return parseListAdminUsersResponse(statusCode, body)
}

func (letmein *OrganizationRepo) AddAdminUser(orgID, email string) (*AdminUserData, error) {
	url := letmein.Repo.BaseURL + "/rest/admin/organizations/" + orgID + "/admin_users"
	req := restapi.New(url, http.MethodPost)
	req.AddHeader("Authorization", fmt.Sprintf("Bearer %s", letmein.Repo.SessionToken))
	statusCode, body, err := req.DoRequest()

	if err != nil {
		return nil, fmt.Errorf("error requesting list of organization's admin users: %s", err)
	}

	return parseAddAdminUserResponse(statusCode, body)
}

func (letmein *OrganizationRepo) RemoveAdminUser(orgID, adminUserID string) error {
	url := letmein.Repo.BaseURL + "/rest/admin/organizations/" + orgID + "/admin_users/" + adminUserID
	req := restapi.New(url, http.MethodPost)
	req.AddHeader("Authorization", fmt.Sprintf("Bearer %s", letmein.Repo.SessionToken))
	statusCode, body, err := req.DoRequest()

	if err != nil {
		return fmt.Errorf("error requesting delete organization's admin users: %s", err)
	}

	return parseRemoveResponse(statusCode, body)
}

func parseListAdminUsersResponse(statusCode int, body []byte) (*AdminUsers, error) {
	var adminUsers *AdminUsers
	var err error
	if statusCode == http.StatusOK {
		adminUsers, err = parseAdminUsersResponse(body)
	} else {
		err = letmeinerr.New(statusCode, body)
	}

	return adminUsers, err
}

func parseAdminUsersResponse(body []byte) (*AdminUsers, error) {
	var adminUsers *AdminUsers
	err := json.Unmarshal(body, &adminUsers)
	if err != nil {
		return nil, fmt.Errorf("json parser error on listing organization's admin users: %w", err)
	}

	return adminUsers, nil
}

func parseAddAdminUserResponse(statusCode int, body []byte) (*AdminUserData, error) {
	var adminUserData *AdminUserData
	var err error
	if statusCode == http.StatusCreated {
		adminUserData, err = parseAdminUserResponse(body)
	} else {
		err = letmeinerr.New(statusCode, body)
	}

	return adminUserData, err
}

func parseAdminUserResponse(body []byte) (*AdminUserData, error) {
	var adminUserData *AdminUserData
	err := json.Unmarshal(body, &adminUserData)
	if err != nil {
		return nil, fmt.Errorf("json parser error on add organization's admin users: %w", err)
	}

	return adminUserData, nil
}

func parseRemoveResponse(statusCode int, body []byte) error {
	var err error

	if statusCode != http.StatusNoContent {
		err = letmeinerr.New(statusCode, body)
	}

	return err
}
