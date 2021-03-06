package cmd

import (
	"io"

	"github.com/notomo/qaper/domain/repository"
	crepository "github.com/notomo/qaper/internal/client/repository"
)

// QuestionCommand represents `question` command
type QuestionCommand struct {
	OutputWriter       io.Writer
	QuestionRepository repository.QuestionRepository
	StateRepository    crepository.StateRepository
}

// Run `question` command
func (c *QuestionCommand) Run() error {
	state, err := c.StateRepository.Load()
	if err != nil {
		return err
	}

	question, err := c.QuestionRepository.GetCurrent(state.PaperID())
	if err != nil {
		return err
	}

	c.OutputWriter.Write([]byte(question.Body()))
	return nil
}
