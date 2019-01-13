package cmd

import (
	"flag"
	"io"
)

// HelpCommand represents `help` command
type HelpCommand struct {
	OutWriter io.Writer
}

// Run `help` command
func (c *HelpCommand) Run() error {
	flag.CommandLine.SetOutput(c.OutWriter)
	flag.Usage()
	return nil
}
