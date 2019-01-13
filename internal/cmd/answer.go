package cmd

import (
	"bufio"
	"io"

	"github.com/notomo/qaper/internal/answer"
)

// AnswerCommand represents `answer` command
type AnswerCommand struct {
	InputReader io.Reader
}

// Run `answer` command
func (c *AnswerCommand) Run() error {
	scan := bufio.NewScanner(c.InputReader)
	if res := scan.Scan(); !res {
		return nil
	}
	ans := scan.Bytes()

	if err := answer.Set(string(ans)); err != nil {
		return err
	}
	return nil
}
