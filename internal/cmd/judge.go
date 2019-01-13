package cmd

import (
	"io"

	"github.com/notomo/qaper/internal/judge"
)

// JudgeCommand represents `judge` command
type JudgeCommand struct {
	OutWriter io.Writer
}

// Run `judge` command
func (c *JudgeCommand) Run() error {
	group, err := judge.Get()
	if err != nil {
		return err
	}
	c.OutWriter.Write([]byte(group.String()))
	return nil
}
