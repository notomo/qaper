package cmd

import (
	"io"

	"github.com/notomo/qaper/internal/question"
)

// QuestionCommand represents `question` command
type QuestionCommand struct {
	OutWriter io.Writer
}

// Run `question` command
func (c *QuestionCommand) Run() error {
	q, err := question.Get()
	if err != nil {
		return err
	}
	c.OutWriter.Write([]byte(q.String()))
	return nil
}
