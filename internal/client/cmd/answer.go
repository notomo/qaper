package cmd

import (
	"io"

	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/domain/repository"
	crepository "github.com/notomo/qaper/internal/client/repository"
)

// AnswerCommand represents `answer` command
type AnswerCommand struct {
	OutputWriter     io.Writer
	StateRepository  crepository.StateRepository
	AnswerRepository repository.AnswerRepository
	Answer           model.Answer
}

// Run `answer` command
func (c *AnswerCommand) Run() error {
	state, err := c.StateRepository.Load()
	if err != nil {
		return err
	}

	if err := c.AnswerRepository.Set(state.PaperID(), c.Answer); err != nil {
		return err
	}

	c.OutputWriter.Write([]byte("ok"))
	return nil
}
