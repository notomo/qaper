package cmd

import (
	"io"

	"github.com/notomo/qaper/domain/repository"
)

// JoinCommand represents `join` command
type JoinCommand struct {
	OutputWriter    io.Writer
	PaperRepository repository.PaperRepository
	BookID          string
}

// Run `join` command
func (c *JoinCommand) Run() error {
	paper, err := c.PaperRepository.Add(c.BookID)
	if err != nil {
		return err
	}

	c.OutputWriter.Write([]byte(paper.ID()))

	return nil
}
