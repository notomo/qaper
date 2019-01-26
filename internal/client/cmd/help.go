package cmd

import (
	"flag"
	"io"
)

// HelpCommand represents `help` command
type HelpCommand struct {
	OutputWriter io.Writer
}

// Run `help` command
func (c *HelpCommand) Run() error {
	flag.CommandLine.SetOutput(c.OutputWriter)
	flag.Usage()
	return nil
}
