package organizations_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adilsonchacon/goeli/app/admin/organizations"
	"github.com/adilsonchacon/goeli/config/admin"
	"github.com/adilsonchacon/goeli/lib/letmeinerr"
)

func TestListAdminUsersSuccess(t *testing.T) {
	jsonResponse := `{
		"data": [
			{
				"id": "123456789",
				"user": {
					"name": "John Doe",
					"email": "john.doe@example.com"
				}
			},
			{
				"id": "382938293",
				"user": {
					"name": "Jane Doe",
					"email": "jane.doe@example.com"
				}
			}
		],
		"pagination": {
			"count": 40,
			"first": 1,
			"last": 2,
			"next": 2,
			"page": 1,
			"per_page": 20,
			"prev": null,
			"serie": [1,2,3,4]
		}
	}`

	sessionToken := "a-valid-token"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonResponse))
	}))

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	adminUsers, err := organizationRepo.ListAdminUsers("123456789", 1, 20)

	if err == nil {
		t.Log("should return list of admin users with pagination")
	} else {
		t.Errorf("List admin users expected, but got error: %s", err)
	}

	if len(adminUsers.Data) == 2 {
		t.Log("List admin users returned expected number of users")
	} else {
		t.Errorf("List admin users expected to return 2 users, got %d", len(adminUsers.Data))
	}

	if adminUsers.Data[0].ID == "123456789" {
		t.Log("First admin user in list has the expected ID")
	} else {
		t.Errorf("First admin user ID expected to be 123456789, got %s", adminUsers.Data[0].ID)
	}

	if adminUsers.Data[1].ID == "382938293" {
		t.Log("Second admin user in list has the expected ID")
	} else {
		t.Errorf("Second admin user ID expected to be 382938293, got %s", adminUsers.Data[0].ID)
	}

	if adminUsers.Pagination.Count == 40 {
		t.Log("Pagination returned expected count")
	} else {
		t.Errorf("Pagination expected to count 40 users, got %d", adminUsers.Pagination.Count)
	}
}

func TestListAdminUsersWithInvalidTokenFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"detail": "Forbidden"
		}
	}`

	sessionToken := "an-invalid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(jsonResponse))
	}))

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	_, err := organizationRepo.ListAdminUsers("123456789", 1, 20)
	if err != nil {
		t.Log("With invalid token should return error")
	} else {
		t.Error("With invalid token did not return error")
	}

	expectedError := &letmeinerr.LetmeinError{
		StatusCode: http.StatusForbidden,
		Body:       []byte(jsonResponse),
		MainError:  letmeinerr.ErrForbidden,
	}

	if errors.As(err, &expectedError) {
		t.Log("With invalid token returns Forbidden")
	} else {
		t.Errorf("With invalid token does not return Forbidden, got %s", err)
	}

}

func TestListAdminUsersWithInvalidOrgIdFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"detail": "Not Found"
		}
	}`

	sessionToken := "a-valid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(jsonResponse))
	}))

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	_, err := organizationRepo.ListAdminUsers("invalid-id", 1, 20)
	if err != nil {
		t.Log("With invalid organization ID should return error")
	} else {
		t.Error("With invalid organization ID not return error")
	}

	expectedError := &letmeinerr.LetmeinError{
		StatusCode: http.StatusNotFound,
		Body:       []byte(jsonResponse),
		MainError:  letmeinerr.ErrNotFound,
	}

	if errors.As(err, &expectedError) {
		t.Log("With invalid organization ID returns Not Found")
	} else {
		t.Errorf("With invalid organization ID returns Not Found, got %s", err)
	}

}
