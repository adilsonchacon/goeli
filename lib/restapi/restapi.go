package restapi

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type RESTApi struct {
	URL        string
	HttpMethod string
	Headers    map[string]string
	Body       map[string]string
}

func New(url string, httpMethod string) *RESTApi {
	return &RESTApi{
		URL:        url,
		HttpMethod: httpMethod,
	}
}

func (restApi *RESTApi) SetHeaders(headers map[string]string) {
	restApi.Headers = headers
}

func (restApi *RESTApi) AddHeader(key, value string) {
	if restApi.Headers == nil {
		restApi.Headers = make(map[string]string)
	}

	restApi.Headers[key] = value
}

func (restApi *RESTApi) SetBody(body map[string]string) {

	restApi.Body = body
}

func (restApi *RESTApi) AddBody(key, value string) {
	if restApi.Body == nil {
		restApi.Body = make(map[string]string)
	}

	restApi.Body[key] = value
}

func (restApi *RESTApi) DoRequest() (int, []byte, error) {
	jsonBody := []byte(restApi.stringyBody())
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(restApi.HttpMethod, restApi.URL, bodyReader)
	if err != nil {
		return 0, nil, fmt.Errorf("could not create request: %s", err)
	}
	req.Header.Set("content-type", "application/json")
	for key, value := range restApi.Headers {
		req.Header.Set(key, scapeString(value))
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("error requesting: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("body reader error: %s", err)
	}

	return res.StatusCode, body, nil
}

func (restApi *RESTApi) stringyBody() string {
	var body []string
	for key, value := range restApi.Body {
		body = append(body, fmt.Sprintf("\"%s\": \"%s\"", key, scapeString(value)))
	}

	return "{" + strings.Join(body, ", ") + "}"
}

func scapeString(str string) string {
	return strings.ReplaceAll(str, "\"", "\\\"")
}
