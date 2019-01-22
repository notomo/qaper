package cmd

import (
	"io"

	"github.com/notomo/qaper/domain/repository"
)

// JoinCommand represents `join` command
type JoinCommand struct {
	OutWriter       io.Writer
	PaperRepository repository.PaperRepository
}

// Run `join` command
func (c *JoinCommand) Run() error {
	paper, err := c.PaperRepository.Add()
	if err != nil {
		return err
	}

	c.OutWriter.Write([]byte(paper.ID()))

	return nil
}
