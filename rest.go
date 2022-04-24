package eventide

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RestError struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte
	Message      *ErrorMessage
}

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("http %d: %s", e.Response.StatusCode, e.ResponseBody)
}

func (c *Client) Request(method string, url string, body io.Reader) ([]byte, error) {
	var err error

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", c.token)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	default:
		e := RestError{
			Request:      req,
			Response:     resp,
			ResponseBody: respBody,
		}

		var mes ErrorMessage
		if err := json.Unmarshal(respBody, &mes); err == nil {
			e.Message = &mes
		}
		err = e
	}

	return respBody, err
}

func (c *Client) GetGatewayURL() (string, error) {
	var err error

	body, err := c.Request("GET", EndpointGateway, nil)
	if err != nil {
		return "", err
	}

	var data GatewayURL
	err = json.Unmarshal(body, &data)

	return data.URL, err
}

func (c *Client) SendMessage(channelID string, message string) (*Message, error) {
	var err error
	payload := &MessageSend{
		Content: message,
	}

	dat, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	body, err := c.Request("POST", EndpointChannelMessages(channelID), bytes.NewBuffer(dat))
	if err != nil {
		return nil, err
	}

	var mes Message
	err = json.Unmarshal(body, &mes)

	return &mes, err
}
