package rchatwrapr

import (
	"bytes"
	"fmt"
	"net/http"
)

func (c *Client) setHeader(req *http.Request) {
	req.Header.Set("X-Auth-Token", c.authToken)
	req.Header.Set("X-User-Id", c.userID)
	req.Header.Set("Content-Type", "application/json")
}
func (c *Client) makeRequest(method, url string, payload []byte) (resp *http.Response, err error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return
	}
	c.setHeader(req)

	resp, err = client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusUnauthorized {
		err = fmt.Errorf("unauthorized to make the request please run init command")
	}
	return
}
