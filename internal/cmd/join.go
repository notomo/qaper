package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// JoinCommand represents `join` command
type JoinCommand struct {
	OutWriter io.Writer
	Port      string
}

// Run `join` command
func (c *JoinCommand) Run() error {
	u := fmt.Sprintf("http://localhost:%s/paper", c.Port)
	res, err := http.PostForm(u, url.Values{})
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	c.OutWriter.Write(body)

	return nil
}
