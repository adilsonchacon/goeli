package organizations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adilsonchacon/goeli/config/admin"
	"github.com/adilsonchacon/goeli/lib/letmeinerr"
	"github.com/adilsonchacon/goeli/lib/restapi"
)

type OrganizationRepo struct {
	Repo *admin.Config
}

func NewRepo(repo *admin.Config) OrganizationRepo {
	return OrganizationRepo{Repo: repo}
}

func (letmein *OrganizationRepo) Create(newOrganization Organization) (*Organization, error) {
	req := restapi.New(letmein.Repo.BaseURL+"/rest/admin/organizations", http.MethodPost)
	req.AddHeader("Authorization", fmt.Sprintf("Bearer %s", letmein.Repo.SessionToken))
	req.AddBody("name", newOrganization.Name)
	req.AddBody("description", newOrganization.Description)
	statusCode, body, err := req.DoRequest()

	if err != nil {
		return nil, fmt.Errorf("error requesting for create organization: %w", err)
	}

	return parseCreateResponse(statusCode, body)
}

func (letmein *OrganizationRepo) Update(organization Organization) (*Organization, error) {
	req := restapi.New(letmein.Repo.BaseURL+"/rest/admin/organizations/"+organization.ID, http.MethodPut)
	req.AddHeader("Authorization", fmt.Sprintf("Bearer %s", letmein.Repo.SessionToken))
	req.AddBody("name", organization.Name)
	req.AddBody("description", organization.Description)
	statusCode, body, err := req.DoRequest()

	if err != nil {
		return nil, fmt.Errorf("error requesting for create organization: %w", err)
	}

	return parseCreateResponse(statusCode, body)
}

func (letmein *OrganizationRepo) Find(id string) (*Organization, error) {
	req := restapi.New(letmein.Repo.BaseURL+"/rest/admin/organizations/"+id, http.MethodGet)
	req.AddHeader("Authorization", fmt.Sprintf("Bearer %s", letmein.Repo.SessionToken))
	statusCode, body, err := req.DoRequest()

	if err != nil {
		return nil, fmt.Errorf("error requesting Find Organization: %s", err)
	}

	return parseFindResponse(statusCode, body)
}

func (letmein *OrganizationRepo) Delete(id string) error {
	req := restapi.New(letmein.Repo.BaseURL+"/rest/admin/organizations/"+id, http.MethodGet)
	req.AddHeader("Authorization", fmt.Sprintf("Bearer %s", letmein.Repo.SessionToken))
	statusCode, body, err := req.DoRequest()

	if err != nil {
		return fmt.Errorf("error requesting delete organization: %s", err)
	}

	return parseDeleteResponse(statusCode, body)
}

func parseCreateResponse(statusCode int, body []byte) (*Organization, error) {
	var organization *Organization
	var err error

	if statusCode == http.StatusCreated {
		organization, err = parseOrganizationResponse(body)
	} else {
		err = letmeinerr.New(statusCode, body)
	}

	return organization, err
}

func parseFindResponse(statusCode int, body []byte) (*Organization, error) {
	var organization *Organization
	var err error

	if statusCode == http.StatusOK {
		organization, err = parseOrganizationResponse(body)
	} else {
		err = letmeinerr.New(statusCode, body)
	}

	return organization, err
}

func parseDeleteResponse(statusCode int, body []byte) error {
	var err error

	if statusCode != http.StatusNoContent {
		err = letmeinerr.New(statusCode, body)
	}

	return err
}

func parseOrganizationResponse(body []byte) (*Organization, error) {
	var dataOrganization *OrganizationData
	err := json.Unmarshal(body, &dataOrganization)
	if err != nil {
		return nil, fmt.Errorf("json parser error on create organization: %w", err)
	}

	return &dataOrganization.Organization, nil
}
