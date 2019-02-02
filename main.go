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
	"github.com/notomo/qaper/internal/datastore"
	"github.com/notomo/qaper/internal/server/api/controller"
	server "github.com/notomo/qaper/internal/server/datastore"
)

type globalConfig struct {
	port string
}

func main() {
	port := flag.String("port", "", "port number")
	conf := &globalConfig{
		port: *port,
	}

	flag.Parse()
	args := flag.Args()

	if err := run(args, conf, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func run(args []string, conf *globalConfig, inputReader io.Reader, outputWriter io.Writer) error {
	command, err := parseCommand(args, conf, inputReader, outputWriter)
	if err != nil {
		return err
	}

	return command.Run()
}

func parseCommand(args []string, conf *globalConfig, inputReader io.Reader, outputWriter io.Writer) (cmd.Command, error) {
	joinFlag := flag.NewFlagSet("join", flag.ExitOnError)
	bookID := joinFlag.String("bookid", "", "book id")

	answerFlag := flag.NewFlagSet("answer", flag.ExitOnError)
	answerBody := answerFlag.String("body", "", "answer body")

	serverFlag := flag.NewFlagSet("server", flag.ExitOnError)
	configPath := serverFlag.String("config", "", "config file path")

	if len(args) == 0 {
		return &cmd.HelpCommand{OutputWriter: outputWriter}, nil
	}

	var command cmd.Command
	switch args[0] {
	case "join":
		if err := joinFlag.Parse(args[1:]); err != nil {
			return nil, err
		}

		joinConfig, err := (&config.JoinConfig{
			Port: conf.port,
		}).Load(*configPath)
		if err != nil {
			return nil, err
		}

		command = &cmd.JoinCommand{
			OutputWriter:    outputWriter,
			PaperRepository: &client.PaperRepositoryImpl{Port: joinConfig.Port},
			StateRepository: &client.StateRepositoryImpl{},
			BookID:          *bookID,
		}
	case "question":
		questionConfig, err := (&config.QuestionConfig{
			Port: conf.port,
		}).Load(*configPath)
		if err != nil {
			return nil, err
		}

		command = &cmd.QuestionCommand{
			OutputWriter:       outputWriter,
			QuestionRepository: &client.QuestionRepositoryImpl{Port: questionConfig.Port},
			StateRepository:    &client.StateRepositoryImpl{},
		}
	case "answer":
		if err := answerFlag.Parse(args[1:]); err != nil {
			return nil, err
		}

		answerConfig, err := (&config.AnswerConfig{
			Port: conf.port,
		}).Load(*configPath)
		if err != nil {
			return nil, err
		}

		command = &cmd.AnswerCommand{
			OutputWriter:     outputWriter,
			AnswerRepository: &client.AnswerRepositoryImpl{Port: answerConfig.Port},
			StateRepository:  &client.StateRepositoryImpl{},
			Answer:           &datastore.AnswerImpl{AnswerBody: *answerBody},
		}
	case "server":
		if err := serverFlag.Parse(args[1:]); err != nil {
			return nil, err
		}

		serverConfig, err := (&config.ServerConfig{
			LibraryPath: "",
			Port:        conf.port,
		}).Load(*configPath)
		if err != nil {
			return nil, err
		}

		processor := server.NewProcessor()
		command = &cmd.ServerCommand{
			OutputWriter: outputWriter,
			Port:         serverConfig.Port,
			LibraryPath:  serverConfig.LibraryPath,
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
