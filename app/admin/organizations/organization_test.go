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

func TestCreateSuccess(t *testing.T) {
	jsonResponse := `{
		"data": {
			"id": "123456789",
			"name": "My Organization",
			"description": "My Organization Description"
		}
	}`

	sessionToken := "a-valid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(jsonResponse))
	}))

	newOrganization := organizations.Organization{
		Name:        "My Organization",
		Description: "My Organization Description",
	}

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	organization, err := organizationRepo.Create(newOrganization)
	if err == nil {
		t.Log("With valid token should create Organization")
	} else {
		t.Errorf("Create Organization expected no errors, got %s", err)
	}

	if organization.ID == "123456789" {
		t.Log("Create Organization returned expected ID")
	} else {
		t.Errorf("Organization ID Expected to be 123456789, got %s", organization.ID)
	}
}

func TestCreateInputFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"name": ["can't be blank"]
		}
	}`

	sessionToken := "a-valid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(jsonResponse))
	}))

	newOrganization := organizations.Organization{
		Name:        "My Organization",
		Description: "My Organization Description",
	}

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	_, err := organizationRepo.Create(newOrganization)

	if err != nil {
		t.Logf("With blank name should return error")
	} else {
		t.Error("With blank name should return errors, got none")
	}

	expectedError := &letmeinerr.LetmeinError{
		StatusCode: http.StatusUnprocessableEntity,
		Body:       []byte(jsonResponse),
		MainError:  letmeinerr.ErrUnprocessableEntity,
	}

	if errors.As(err, &expectedError) {
		t.Logf("With blank name returns UnprocessableEntityError")
	} else {
		t.Errorf("With blank name does not return UnprocessableEntityError, got %s", err)
	}
}

func TestCreateTokenFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"detail": "Invalid token"
		}
	}`

	sessionToken := "an-invalid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(jsonResponse))
	}))

	newOrganization := organizations.Organization{
		Name:        "My Organization",
		Description: "My Organization Description",
	}

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	_, err := organizationRepo.Create(newOrganization)
	if err != nil {
		t.Log("With invalid token should return error")
	} else {
		t.Error("With invalid token should return error")
	}

	expectedError := &letmeinerr.LetmeinError{
		StatusCode: http.StatusForbidden,
		Body:       []byte(jsonResponse),
		MainError:  letmeinerr.ErrForbidden,
	}

	if errors.As(err, &expectedError) {
		t.Logf("With invalid token returns ForbiddenError")
	} else {
		t.Errorf("With invalid token does not return ForbiddenError, got %s", err)
	}
}

func TestFindSuccess(t *testing.T) {
	jsonResponse := `{
		"data": {
			"id": "123456789",
			"name": "My Organization",
			"description": "My Organization Description"
		}
	}`

	sessionToken := "a-valid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonResponse))
	}))

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	organization, err := organizationRepo.Find("123456789")
	if err == nil {
		t.Log("With existent ID should return Organization")
	} else {
		t.Errorf("Find Organization expected no errors, got %s", err)
	}

	if organization.ID == "123456789" {
		t.Log("Find Organization returned expected ID")
	} else {
		t.Errorf("Organization ID Expected to be 123456789, got %s", organization.ID)
	}
}

func TestFindWithNotExistentIdFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"detail": "Not found"
		}
	}`

	sessionToken := "a-valid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(jsonResponse))
	}))

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	_, err := organizationRepo.Find("987654321")
	if err != nil {
		t.Log("With invalid ID should return error")
	} else {
		t.Error("With invalid ID did not return error")
	}

	expectedError := &letmeinerr.LetmeinError{
		StatusCode: http.StatusNotFound,
		Body:       []byte(jsonResponse),
		MainError:  letmeinerr.ErrNotFound,
	}

	if errors.As(err, &expectedError) {
		t.Log("With a not existent ID returns NotFoundError")
	} else {
		t.Errorf("With a not existent ID does not return NotFoundError, got %s", err)
	}
}

func TestFindWithInvalidTokenFails(t *testing.T) {
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
	_, err := organizationRepo.Find("987654321")
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

func TestUpdateSuccess(t *testing.T) {
	jsonResponse := `{
		"data": {
			"id": "123456789",
			"name": "Updated Organization Name",
			"description": "Updated Organization Description"
		}
	}`

	sessionToken := "a-valid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(jsonResponse))
	}))

	updateOrganization := organizations.Organization{
		Name:        "Updated Organization Name",
		Description: "UpdatedUpdated Organization Description",
	}

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	_, err := organizationRepo.Update(updateOrganization)
	if err == nil {
		t.Log("With valid token should update Organization")
	} else {
		t.Errorf("Update Organization expected no errors, got %s", err)
	}
}

func TestUpdateInputFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"name": ["can't be blank"]
		}
	}`

	sessionToken := "a-valid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(jsonResponse))
	}))

	updateOrganization := organizations.Organization{
		Name:        "",
		Description: "UpdatedUpdated Organization Description",
	}

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	_, err := organizationRepo.Update(updateOrganization)
	if err != nil {
		t.Logf("With blank name should return error")
	} else {
		t.Error("With blank name should return errors, got none")
	}

	expectedError := &letmeinerr.LetmeinError{
		StatusCode: http.StatusUnprocessableEntity,
		Body:       []byte(jsonResponse),
		MainError:  letmeinerr.ErrUnprocessableEntity,
	}

	if errors.As(err, &expectedError) {
		t.Logf("With blank name returns UnprocessableEntityError")
	} else {
		t.Errorf("With blank name does not return UnprocessableEntityError, got %s", err)
	}
}

func TestUpdateTokenFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"detail": "Invalid token"
		}
	}`

	sessionToken := "an-invalid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(jsonResponse))
	}))

	updateOrganization := organizations.Organization{
		Name:        "Updated Organization Name",
		Description: "UpdatedUpdated Organization Description",
	}

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	_, err := organizationRepo.Update(updateOrganization)
	if err != nil {
		t.Log("With invalid token should return error")
	} else {
		t.Error("With invalid token should return error")
	}

	expectedError := &letmeinerr.LetmeinError{
		StatusCode: http.StatusForbidden,
		Body:       []byte(jsonResponse),
		MainError:  letmeinerr.ErrForbidden,
	}

	if errors.As(err, &expectedError) {
		t.Logf("With invalid token returns ForbiddenError")
	} else {
		t.Errorf("With invalid token does not return ForbiddenError, got %s", err)
	}
}

func TestDeleteSuccess(t *testing.T) {
	jsonResponse := ``

	sessionToken := "a-valid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(jsonResponse))
	}))

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	err := organizationRepo.Delete("123456789")
	if err == nil {
		t.Log("With valid token should delete Organization")
	} else {
		t.Errorf("Delete Organization expected no errors, got %s", err)
	}
}

func TestDeleteInputFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"details": "Not Found"
		}
	}`

	sessionToken := "a-valid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(jsonResponse))
	}))

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	err := organizationRepo.Delete("123456789")
	if err != nil {
		t.Logf("With not existent ID should return error")
	} else {
		t.Error("With not existent ID should return errors, got none")
	}

	expectedError := &letmeinerr.LetmeinError{
		StatusCode: http.StatusNotFound,
		Body:       []byte(jsonResponse),
		MainError:  letmeinerr.ErrNotFound,
	}

	if errors.As(err, &expectedError) {
		t.Logf("With not existent ID returns ErrNotFound")
	} else {
		t.Errorf("With not existent ID does not return ErrNotFound, got %s", err)
	}
}

func TestDeleteTokenFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"detail": "Invalid token"
		}
	}`

	sessionToken := "an-invalid-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(jsonResponse))
	}))

	adminConfig := admin.NewConfig(server.URL, sessionToken)
	organizationRepo := organizations.NewRepo(adminConfig)
	err := organizationRepo.Delete("123456789")
	if err != nil {
		t.Log("With invalid token should return error")
	} else {
		t.Error("With invalid token should return error")
	}

	expectedError := &letmeinerr.LetmeinError{
		StatusCode: http.StatusNotFound,
		Body:       []byte(jsonResponse),
		MainError:  letmeinerr.ErrNotFound,
	}

	if errors.As(err, &expectedError) {
		t.Logf("With invalid token returns ForbiddenError")
	} else {
		t.Errorf("With invalid token does not return ForbiddenError, got %s", err)
	}
}
