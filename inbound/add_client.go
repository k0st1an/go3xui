package inbound

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/k0st1an/go3xui"
)

type addClient struct {
	ID       int64  `json:"id"` // ID of the inbound
	Settings string `json:"settings"`
}

type AddClient struct {
	ID       int64             `json:"id"` // ID of the inbound
	Settings AddClientSettings `json:"settings"`
}

type AddClientSettings struct {
	Client []AddClientClient `json:"clients"`
}

type AddClientClient struct {
	ID         string `json:"id"`    // example: 4f83fe3a-6f3a-4311-b92d-898b31e9a0f0
	Flow       string `json:"flow"`  // example: xtls-rprx-vision
	Email      string `json:"email"` // example: on40d6bg
	LimitIP    int    `json:"limitIp"`
	TotalGB    int    `json:"totalGB"`
	ExpiryTime int    `json:"expiryTime"` // example: 1750881979990
	Enable     bool   `json:"enable"`
	TgId       string `json:"tgId"`
	SubId      string `json:"subId"`   // Subscription: 66jnhbckta7kfxyz
	Comment    string `json:"comment"` // example: 222111
	Reset      int    `json:"reset"`
}

type AddClientResp struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Obj     any    `json:"obj"`
}

func (a *AddClient) AddClient(c *go3xui.Client) (*AddClientResp, error) {
	settings, err := json.Marshal(a.Settings)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal settings: %w", err)
	}

	addNewClient := addClient{
		ID:       a.ID,
		Settings: string(settings),
	}

	jsonData, err := json.Marshal(addNewClient)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", go3xui.EndpointInboundAddClient(c.URL), bytes.NewBuffer(jsonData))
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

	var addClientResp AddClientResp
	if err := json.NewDecoder(resp.Body).Decode(&addClientResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if !addClientResp.Success {
		return nil, fmt.Errorf("failed to add client: %s", addClientResp.Msg)
	}

	return &addClientResp, nil
}
