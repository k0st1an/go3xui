package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/k0st1an/go3xui"
)

func GetServerStatus(c *go3xui.Client) (*ServerStatusResp, error) {
	url := go3xui.EndpointServerStatus(c.URL)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var serverStatus ServerStatusResp
	if err := json.Unmarshal(body, &serverStatus); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &serverStatus, nil
}
