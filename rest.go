package eventide

import (
	"net/http"
	"encoding/json"
)

func (c *Client) Request(method string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if c.Token != "" {
		req.Header.Set("Authorization", c.Token)
	}

	resp, err := c.http.Do(req)

	return resp, err
}

func (c *Client) GetGatewayURL() (string, error) {
	var err error

	resp, err := c.Request("GET", "https://discord.com/api/v9/gateway")
	if err != nil {
		return "", err
	}

	var data struct {
		Url string `json:"url"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)

	return data.Url, err
}