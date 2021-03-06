package cmd

import (
	"io"

	"github.com/notomo/qaper/domain/repository"
	crepository "github.com/notomo/qaper/internal/client/repository"
)

// JoinCommand represents `join` command
type JoinCommand struct {
	OutputWriter    io.Writer
	PaperRepository repository.PaperRepository
	StateRepository crepository.StateRepository
	BookID          string
}

// Run `join` command
func (c *JoinCommand) Run() error {
	paper, err := c.PaperRepository.Add(c.BookID)
	if err != nil {
		return err
	}

	if err := c.StateRepository.Save(paper); err != nil {
		return err
	}

	c.OutputWriter.Write([]byte(paper.ID()))

	return nil
}
