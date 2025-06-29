package admin

type Config struct {
	BaseURL      string
	SessionToken string
}

func NewConfig(baseURL, sessionToken string) *Config {
	return &Config{
		BaseURL:      baseURL,
		SessionToken: sessionToken,
	}
}
