package go3xui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Login(username, password string) error {
	loginData := LoginReq{
		Username: username,
		Password: password,
		// LoginSecret: loginSecret,	// TODO: implement
	}

	jsonData, err := json.Marshal(loginData)
	if err != nil {
		return fmt.Errorf("failed to marshal login data: %w", err)
	}

	req, err := http.NewRequest("POST", EndpointLogin(c.URL), bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create login request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("login request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("login failed with status %d: %s", resp.StatusCode, string(body))
	}

	var loginResponse LoginResp
	if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		return fmt.Errorf("failed to decode login response: %w", err)
	}

	if !loginResponse.Success {
		return fmt.Errorf("login failed: %s", loginResponse.Msg)
	}

	return nil
}

type LoginReq struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	LoginSecret string `json:"loginSecret"`
}

type LoginResp struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Obj     any    `json:"obj"`
}
