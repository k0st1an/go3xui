package go3xui

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	HttpClient *http.Client
	URL        string
}

func NewClient(url string) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar: %w", err)
	}

	client := &http.Client{Jar: jar}

	return &Client{
		HttpClient: client,
		URL:        url,
	}, nil
}
