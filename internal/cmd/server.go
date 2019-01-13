package cmd

import (
	"io"

	"github.com/notomo/qaper/internal/server"
)

// ServerCommand represents `server` command
type ServerCommand struct {
	OutWriter io.Writer
	Port      string
}

// Run `server` command
func (c *ServerCommand) Run() error {
	config := server.Config{Port: c.Port}
	s := config.Server()
	return s.Listen()
}
