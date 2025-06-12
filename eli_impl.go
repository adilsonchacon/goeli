package goeli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/adilsonchacon/goeli/entities"
)

func (config *Config) SignIn(email, password string) (string, int, error) {
	requestURL := config.BaseURL + "/rest" + addAdminToUrlPath(config.ServiceType) + "/sessions"
	jsonBody := []byte(`{"email": "` + email + `", "password": "` + password + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return "", 0, fmt.Errorf("could not create request for SignIn: %s", err)
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("app-token", config.AppToken)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("error requesting for SignIn: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", 0, fmt.Errorf("body reader error for SignIn: %s", err)
	}

	return parseSignInResponse(res.StatusCode, body)
}

func (config *Config) SignedIn(sessionToken string) (bool, error) {
	requestURL := config.BaseURL + "/rest" + addAdminToUrlPath(config.ServiceType) + "/sessions/signed_in"

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return false, fmt.Errorf("could not create request for SignedIn: %s", err)
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", "Bearer "+sessionToken)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("error requesting for SignedIn: %s", err)
	}
	defer res.Body.Close()

	return res.StatusCode == 200, nil
}

func (config *Config) CurrentUser(sessionToken string) (*entities.User, int, error) {
	requestURL := config.BaseURL + "/rest" + addAdminToUrlPath(config.ServiceType) + "/sessions"

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("could not create request for CurrentUser: %s", err)
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", "Bearer "+sessionToken)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("error requesting for CurrentUser: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("body reader error for SignIn: %s", err)
	}

	return parseCurrentUserResponse(res.StatusCode, body)
}

func (config *Config) SignOut(sessionToken string) (int, error) {
	requestURL := config.BaseURL + "/rest" + addAdminToUrlPath(config.ServiceType) + "/sessions"

	req, err := http.NewRequest(http.MethodDelete, requestURL, nil)
	if err != nil {
		return 0, fmt.Errorf("could not create request for SignOut: %s", err)
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", "Bearer "+sessionToken)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error requesting for SignOut: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("body reader error for SignOut: %s", err)
	}

	return parseSignOutResponse(res.StatusCode, body)
}

func (config *Config) Refresh(sessionToken string) (string, int, error) {
	requestURL := config.BaseURL + "/rest" + addAdminToUrlPath(config.ServiceType) + "/sessions"

	req, err := http.NewRequest(http.MethodPut, requestURL, nil)
	if err != nil {
		return "", 0, fmt.Errorf("could not create request for Refresh: %s", err)
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", "Bearer "+sessionToken)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("error requesting for Refresh: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", 0, fmt.Errorf("body reader error for Refresh: %s", err)
	}

	return parseRefreshResponse(res.StatusCode, body)
}

func (config *Config) Unlock(unlockToken string) (int, error) {
	requestURL := config.BaseURL + "/rest" + addAdminToUrlPath(config.ServiceType) + "/accounts/unlock"
	jsonBody := []byte(`{"token": "` + unlockToken + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPut, requestURL, bodyReader)
	if err != nil {
		return 0, fmt.Errorf("could not create request for Refresh: %s", err)
	}
	req.Header.Set("content-type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error requesting for Unlock: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("body reader error for Unlock: %s", err)
	}

	return parseDefaultAccountResponse(res.StatusCode, body)
}

func (config *Config) Confirm(confirmationToken string) (int, error) {
	requestURL := config.BaseURL + "/rest" + addAdminToUrlPath(config.ServiceType) + "/accounts/confirm"
	jsonBody := []byte(`{"token": "` + confirmationToken + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPut, requestURL, bodyReader)
	if err != nil {
		return 0, fmt.Errorf("could not create request for Refresh: %s", err)
	}
	req.Header.Set("content-type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error requesting for Unlock: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("body reader error for Unlock: %s", err)
	}

	return parseDefaultAccountResponse(res.StatusCode, body)
}

func (config *Config) RequestPasswordRecovery(appToken, email string) (int, error) {
	requestURL := config.BaseURL + "/rest" + addAdminToUrlPath(config.ServiceType) + "/accounts/password/recover"
	jsonBody := []byte(`{"email": "` + email + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return 0, fmt.Errorf("could not create request for Refresh: %s", err)
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("app-token", appToken)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error requesting for Unlock: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("body reader error for Unlock: %s", err)
	}

	return parseRequestPasswordRecoveryResponse(res.StatusCode, body)
}

func (config *Config) RecoverPassword(token, password, passwordConfirmation string) (int, error) {
	return 0, errors.New("some error")
}

func addAdminToUrlPath(serviceName string) string {
	if serviceName == "admin" {
		return "/admin"
	}
	return ""
}

func parseSignInResponse(statusCode int, body []byte) (string, int, error) {
	var message string
	var err error

	if statusCode == http.StatusOK {
		message, err = parseTokenResponse(body)
	} else {
		message, err = parseErrorResponse(body)
	}

	return message, statusCode, err
}

func parseTokenResponse(body []byte) (string, error) {
	var token *entities.LetmeinToken
	err := json.Unmarshal(body, &token)
	if err != nil {
		return "", fmt.Errorf("json parser error: %s", err)
	}

	return token.Data.Token, nil
}

func parseErrorResponse(body []byte) (string, error) {
	var letmeinError *entities.LetmeinError
	err := json.Unmarshal(body, &letmeinError)
	if err != nil {
		return "", fmt.Errorf("json parser error for SignIn when status ERROR: %s", err)
	}

	return "", errors.New(letmeinError.Errors.Detail)
}

func parseUserResponse(body []byte) (*entities.User, int, error) {
	var dataUser *entities.DataUser
	err := json.Unmarshal(body, &dataUser)
	if err != nil {
		return nil, 0, fmt.Errorf("json parser error: %s", err)
	}

	return &dataUser.User, 200, nil
}

func parseCurrentUserResponse(statusCode int, body []byte) (*entities.User, int, error) {
	if statusCode == http.StatusOK {
		return parseUserResponse(body)
	} else {
		_, err := parseErrorResponse(body)
		return nil, statusCode, err
	}
}

func parseSignOutResponse(statusCode int, body []byte) (int, error) {
	if statusCode == http.StatusOK {
		return statusCode, nil
	} else {
		_, err := parseErrorResponse(body)
		return statusCode, err
	}
}

func parseRefreshResponse(statusCode int, body []byte) (string, int, error) {
	var message string
	var err error

	if statusCode == http.StatusOK {
		message, err = parseTokenResponse(body)
	} else {
		message, err = parseErrorResponse(body)
	}

	return message, statusCode, err
}

func parseDefaultAccountResponse(statusCode int, body []byte) (int, error) {
	if statusCode == http.StatusAccepted {
		return statusCode, nil
	} else {
		_, err := parseErrorResponse(body)
		return statusCode, err
	}
}

func parseRequestPasswordRecoveryResponse(statusCode int, body []byte) (int, error) {
	if statusCode == http.StatusOK {
		return statusCode, nil
	} else {
		_, err := parseErrorResponse(body)
		return statusCode, err
	}
}
