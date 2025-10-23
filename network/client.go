// network/client.go
package network

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Client struct {
	http *http.Client
}

func NewClient(timeout time.Duration) *Client {
	return &Client{http: &http.Client{Timeout: timeout}}
}

func (c *Client) Send(method, url string, body string, headers map[string]string) (string, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	var pretty bytes.Buffer
	if json.Valid(b) {
		if err := json.Indent(&pretty, b, "", ""); err == nil {
			return resp.Status + "\n\n" + pretty.String(), nil
		}
	}
	return resp.Status + "\n\n" + string(b), nil
}
