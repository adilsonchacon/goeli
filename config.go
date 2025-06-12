package goeli

import "strings"

type Config struct {
	ServiceType string
	BaseURL     string
	AppToken    string
}

func NewServiceConfig(serviceType, baseURL, appToken string) *Config {
	return &Config{
		ServiceType: normalizeServiceType(serviceType),
		BaseURL:     baseURL,
		AppToken:    appToken,
	}
}

func normalizeServiceType(serviceType string) string {
	if strings.ToLower(serviceType) == "admin" {
		return "admin"
	}

	return "regular"
}
