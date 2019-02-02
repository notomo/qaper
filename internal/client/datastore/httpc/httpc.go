package httpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client represents http client
type Client struct {
	Port   string
	Domain string
}

// Post requests
func (c *Client) Post(path string, v interface{}) (*http.Response, []byte, error) {
	content, err := json.Marshal(v)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("http://%s:%s%s", c.Domain, c.Port, path)
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(content))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	return res, body, nil
}

// Get requests
func (c *Client) Get(path string) (*http.Response, []byte, error) {
	u := fmt.Sprintf("http://%s:%s%s", c.Domain, c.Port, path)
	res, err := http.Get(u)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	return res, body, nil
}
