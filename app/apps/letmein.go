package apps

type Letmein struct {
	BaseURL      string
	SessionToken string
}

func NewLetmein(baseURL, sessionToken string) *Letmein {
	return &Letmein{
		BaseURL:      baseURL,
		SessionToken: sessionToken,
	}
}
