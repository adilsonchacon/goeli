package goeli_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adilsonchacon/goeli"
)

func TestSignInSuccess(t *testing.T) {
	jsonResponse := `{
		"data": {
			"token": "a-valid-token"
		}
 	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	eli := goeli.NewServiceConfig("", server.URL, "some-app-token")

	message, statusCode, err := eli.SignIn("test@test.com", "Secret.123!")
	if err == nil {
		t.Log("[PASSED] with valid credentials SignIn returns no error")
	} else {
		t.Error("[FAILED] with valid credentials SignIn returned an error")
	}

	if statusCode == 200 {
		t.Log("[PASSED] with valid credentials SignIn returns status code 200")
	} else {
		t.Errorf("[FAILED] with valid credentials SignIn did not return status code 200, but %d", statusCode)
	}

	if message == "a-valid-token" {
		t.Log("[PASSED] with valid credentials SignIn returns the token")
	} else {
		t.Errorf("[FAILED] with valid credentials SignIn did not return the token, but %s", message)
	}
}

func TestSignInFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"detail": "invalid credentials"
		}
 	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	user := goeli.NewServiceConfig("", server.URL, "some-app-token")

	message, statusCode, err := user.SignIn("test@test.com", "wrong-password")
	if err != nil {
		t.Log("[PASSED] with invalid credentials SignIn returns an error")
	} else {
		t.Error("[FAILED] with invalid credentials SignIn returns no error")
	}

	if statusCode == 400 {
		t.Log("[PASSED] with invalid credentials SignIn returns status code 400")
	} else {
		t.Errorf("[FAILED] with invalid credentials SignIn did not return status code 400, but %d", statusCode)
	}

	if message == "" {
		t.Log("[PASSED] with invalid credentials SignIn returns an empty message")
	} else {
		t.Errorf("[FAILED] with invalid credentials SignIn did not returned message: %s", message)
	}
}

func TestSignedInSuccess(t *testing.T) {
	jsonResponse := `{}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	user := goeli.NewServiceConfig("", server.URL, "some-app-token")

	signedIn, _ := user.SignedIn("a-valid-token")
	if signedIn {
		t.Log("[PASSED] with valid token SignedIn returns true")
	} else {
		t.Error("[FAILED] with valid token SignedIn returns false")
	}
}

func TestSignedInFails(t *testing.T) {
	jsonResponse := `{}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	user := goeli.NewServiceConfig("", server.URL, "some-app-token")

	signedIn, _ := user.SignedIn("an-invalid-token")
	if !signedIn {
		t.Log("[PASSED] with an invalid token SignedIn returns false")
	} else {
		t.Error("[FAILED] with an invalid token SignedIn returns true")
	}
}

func TestCurrentUserSuccess(t *testing.T) {
	jsonResponse := `{
		"data": {
			"id": "1",
			"email": "test@test.com",
			"name": "Test",
			"active": true,
			"language": "en",
			"timezone": "Europe/London"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	user := goeli.NewServiceConfig("", server.URL, "some-app-token")

	currentUser, statusCode, err := user.CurrentUser("a-valid-token")
	if err == nil {
		t.Log("[PASSED] with a valid token CurrentUser returns no error")
	} else {
		t.Errorf("[FAILED] with a valid token CurrentUser returned error: %s", err)
	}

	if currentUser.ID == "1" {
		t.Log("[PASSED] with a valid token CurrentUser returns an user with id 1")
	} else {
		t.Error("[FAILED] with a valid token CurrentUser should return user with id 1")
	}

	if statusCode == 200 {
		t.Log("[PASSED] with a valid token CurrentUser returns status code 200")
	} else {
		t.Errorf("[FAILED] with a valid token CurrentUser should return status code 200, but returned %d", statusCode)
	}
}

func TestCurrentUserFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"detail": "Not found"
		}
 	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	user := goeli.NewServiceConfig("", server.URL, "some-app-token")

	currentUser, statusCode, err := user.CurrentUser("an-invalid-token")
	if err != nil {
		t.Log("[PASSED] with an invalid token CurrentUser returns error")
	} else {
		t.Error("[FAILED] with an invalid token CurrentUser did not returned error")
	}

	if currentUser == nil {
		t.Log("[PASSED] with an invalid token CurrentUser does not return user data")
	} else {
		t.Error("[FAILED] with an invalid token CurrentUser should not return user data")
	}

	if statusCode != 200 {
		t.Log("[PASSED] with an invalid token CurrentUser does not return status code 200")
	} else {
		t.Errorf("[FAILED] with an invalid token CurrentUser returned status code 200")
	}
}

func TestSignOutSuccess(t *testing.T) {
	jsonResponse := `{
		"data": {
			"message": "signed out successfully"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	user := goeli.NewServiceConfig("", server.URL, "some-app-token")

	statusCode, err := user.SignOut("a-valid-token")
	if err == nil {
		t.Log("[PASSED] with a valid token SignOut returns no error")
	} else {
		t.Errorf("[FAILED] with a valid token SignOut returned error: %s", err)
	}

	if statusCode == 200 {
		t.Log("[PASSED] with a valid token SignOut returns status code 200")
	} else {
		t.Errorf("[FAILED] with a valid token SignOut should return status code 200, but returned %d", statusCode)
	}
}

func TestSignOutFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"detail": "Not found"
		}
 	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	user := goeli.NewServiceConfig("", server.URL, "some-app-token")

	statusCode, err := user.SignOut("an-invalid-token")
	if err != nil {
		t.Log("[PASSED] with an invalid token SignOut returns error")
	} else {
		t.Error("[FAILED] with an invalid token SignOut did not returned error")
	}

	if statusCode != 200 {
		t.Log("[PASSED] with an invalid token SignOut does not return status code 200")
	} else {
		t.Errorf("[FAILED] with an invalid token SignOut returned status code 200")
	}
}

func TestRefreshSuccess(t *testing.T) {
	jsonResponse := `{
		"data": {
			"token": "a-new-and-valid-token"
		}
 	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	user := goeli.NewServiceConfig("", server.URL, "some-app-token")

	message, statusCode, err := user.Refresh("a-valid-token")
	if err == nil {
		t.Log("[PASSED] with a valid token Refresh returns no error")
	} else {
		t.Error("[FAILED] with a valid token Refresh returned an error")
	}

	if statusCode == 200 {
		t.Log("[PASSED] with a valid token Refresh returns status code 200")
	} else {
		t.Errorf("[FAILED] with a valid token Refresh did not return status code 200, but %d", statusCode)
	}

	if message == "a-new-and-valid-token" {
		t.Log("[PASSED] with a valid token Refresh returns the token")
	} else {
		t.Errorf("[FAILED] with a valid token Refresh did not return the token, but %s", message)
	}
}

func TestRefreshFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"detail": "invalid credentials"
		}
 	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	user := goeli.NewServiceConfig("", server.URL, "some-app-token")

	message, statusCode, err := user.Refresh("an-invalid-token")
	if err != nil {
		t.Log("[PASSED] with an invalid token Refresh returns an error")
	} else {
		t.Error("[FAILED] with an invalid token Refresh returns no error")
	}

	if statusCode == 400 {
		t.Log("[PASSED] with an invalid token Refresh returns status code 400")
	} else {
		t.Errorf("[FAILED] with an invalid token Refresh did not return status code 400, but %d", statusCode)
	}

	if message == "" {
		t.Log("[PASSED] with an invalid token Refresh returns an empty message")
	} else {
		t.Errorf("[FAILED] with an invalid token Refresh did not returned message: %s", message)
	}
}

func TestUnlockSuccess(t *testing.T) {
	jsonResponse := `{
		"data": {
			"token": "a-new-and-valid-token"
		}
 	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(202)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	user := goeli.NewServiceConfig("", server.URL, "some-app-token")

	statusCode, err := user.Unlock("a-valid-token")
	if err == nil {
		t.Log("[PASSED] with a valid token Unlock returns no error")
	} else {
		t.Error("[FAILED] with a valid token Unlock returned an error")
	}

	if statusCode == 202 {
		t.Log("[PASSED] with a valid token Unlock returns status code 202")
	} else {
		t.Errorf("[FAILED] with a valid token Unlock did not return status code 202, but %d", statusCode)
	}
}

func TestUnlockFails(t *testing.T) {
	jsonResponse := `{
		"errors": {
			"detail": "invalid credentials"
		}
 	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	user := goeli.NewServiceConfig("", server.URL, "some-app-token")

	statusCode, err := user.Unlock("an-invalid-token")
	if err != nil {
		t.Log("[PASSED] with an invalid token Unlock returns an error")
	} else {
		t.Error("[FAILED] with an invalid token Unlock returns no error")
	}

	if statusCode == 404 {
		t.Log("[PASSED] with an invalid token Unlock returns status code 404")
	} else {
		t.Errorf("[FAILED] with an invalid token Unlock did not return status code 404, but %d", statusCode)
	}
}
