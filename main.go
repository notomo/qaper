package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/notomo/qaper/internal/client/cmd"
	"github.com/notomo/qaper/internal/client/cmd/config"
	client "github.com/notomo/qaper/internal/client/datastore"
	"github.com/notomo/qaper/internal/client/datastore/httpc"
	"github.com/notomo/qaper/internal/datastore"
	"github.com/notomo/qaper/internal/server/api/controller"
	server "github.com/notomo/qaper/internal/server/datastore"
)

var domain = "localhost"

func main() {
	port := flag.String("port", "", "port number")
	configPath := flag.String("config", "", "config file path")
	flag.Parse()

	conf := &config.Config{
		Port:       *port,
		ConfigPath: *configPath,
	}

	args := flag.Args()
	if err := run(args, conf, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func run(args []string, conf *config.Config, inputReader io.Reader, outputWriter io.Writer) error {
	command, err := parseCommand(args, conf, inputReader, outputWriter)
	if err != nil {
		return err
	}

	return command.Run()
}

func parseCommand(args []string, conf *config.Config, inputReader io.Reader, outputWriter io.Writer) (cmd.Command, error) {
	joinFlag := flag.NewFlagSet("join", flag.ExitOnError)
	bookID := joinFlag.String("bookid", "", "book id")

	answerFlag := flag.NewFlagSet("answer", flag.ExitOnError)
	answerBody := answerFlag.String("body", "", "answer body")

	if err := conf.Load(); err != nil {
		return nil, err
	}

	if len(args) == 0 {
		return &cmd.HelpCommand{OutputWriter: outputWriter}, nil
	}

	var command cmd.Command
	switch args[0] {
	case "join":
		if err := joinFlag.Parse(args[1:]); err != nil {
			return nil, err
		}
		command = &cmd.JoinCommand{
			OutputWriter: outputWriter,
			PaperRepository: &client.PaperRepositoryImpl{
				Client: &httpc.Client{
					Port:   conf.Port,
					Domain: domain,
				},
			},
			StateRepository: &client.StateRepositoryImpl{},
			BookID:          *bookID,
		}

	case "question":
		command = &cmd.QuestionCommand{
			OutputWriter: outputWriter,
			QuestionRepository: &client.QuestionRepositoryImpl{
				Client: &httpc.Client{
					Port:   conf.Port,
					Domain: domain,
				},
			},
			StateRepository: &client.StateRepositoryImpl{},
		}

	case "answer":
		if err := answerFlag.Parse(args[1:]); err != nil {
			return nil, err
		}
		command = &cmd.AnswerCommand{
			OutputWriter: outputWriter,
			AnswerRepository: &client.AnswerRepositoryImpl{
				Client: &httpc.Client{
					Port:   conf.Port,
					Domain: domain,
				},
			},
			StateRepository: &client.StateRepositoryImpl{},
			Answer:          &datastore.AnswerImpl{AnswerBody: *answerBody},
		}

	case "server":
		processor := server.NewProcessor()
		command = &cmd.ServerCommand{
			OutputWriter: outputWriter,
			Port:         conf.Port,
			LibraryPath:  conf.Server.LibraryPath,
			Processor:    processor,
			PaperController: controller.PaperController{
				PaperRepository: &server.PaperRepositoryImpl{
					Processor: processor,
					BookRepository: &server.BookRepositoryImpl{
						Processor: processor,
					},
				},
				AnswerRepository: &server.AnswerRepositoryImpl{
					Processor: processor,
				},
				AnswerDecoder: &datastore.AnswerJSONDecoder{},
			},
		}

	default:
		return nil, fmt.Errorf("Not found command: %v", args[0])
	}

	return command, nil
}
