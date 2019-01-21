package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/notomo/qaper/internal/cmd"
	"github.com/notomo/qaper/internal/datastore"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if err := run(args, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func run(args []string, inputReader io.Reader, outWriter io.Writer) error {
	command, err := parseCommand(args, inputReader, outWriter)
	if err != nil {
		return err
	}

	if err := command.Run(); err != nil {
		return err
	}

	return nil
}

func parseCommand(args []string, inputReader io.Reader, outWriter io.Writer) (cmd.Command, error) {
	join := flag.NewFlagSet("join", flag.ExitOnError)
	joinPort := join.String("port", "9090", "port number")

	question := flag.NewFlagSet("question", flag.ExitOnError)
	answer := flag.NewFlagSet("answer", flag.ExitOnError)
	judge := flag.NewFlagSet("judge", flag.ExitOnError)

	server := flag.NewFlagSet("server", flag.ExitOnError)
	serverPort := server.String("port", "9090", "port number")

	if len(args) == 0 {
		return &cmd.HelpCommand{OutWriter: outWriter}, nil
	}

	switch args[0] {
	case "join":
		if err := join.Parse(args[1:]); err != nil {
			return nil, err
		}

		paperRepository := &datastore.PaperRepositoryImpl{Port: *joinPort}
		return &cmd.JoinCommand{OutWriter: outWriter, PaperRepository: paperRepository}, nil
	case "question":
		if err := question.Parse(args[1:]); err != nil {
			return nil, err
		}
		return &cmd.QuestionCommand{OutWriter: outWriter}, nil
	case "answer":
		if err := answer.Parse(args[1:]); err != nil {
			return nil, err
		}
		return &cmd.AnswerCommand{InputReader: inputReader}, nil
	case "judge":
		if err := judge.Parse(args[1:]); err != nil {
			return nil, err
		}
		return &cmd.JudgeCommand{OutWriter: outWriter}, nil
	case "server":
		if err := server.Parse(args[1:]); err != nil {
			return nil, err
		}
		return &cmd.ServerCommand{OutWriter: outWriter, Port: *serverPort}, nil
	default:
		return nil, fmt.Errorf("Not found command: %v", args[0])
	}
}
