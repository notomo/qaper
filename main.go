package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/notomo/qaper/internal/client/cmd"
	client "github.com/notomo/qaper/internal/client/datastore"
	"github.com/notomo/qaper/internal/server/api/controller"
	server "github.com/notomo/qaper/internal/server/datastore"
)

var defaultPort = "9090"

func main() {
	flag.Parse()
	args := flag.Args()

	if err := run(args, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func run(args []string, inputReader io.Reader, outputWriter io.Writer) error {
	command, err := parseCommand(args, inputReader, outputWriter)
	if err != nil {
		return err
	}

	return command.Run()
}

func parseCommand(args []string, inputReader io.Reader, outputWriter io.Writer) (cmd.Command, error) {
	joinFlag := flag.NewFlagSet("join", flag.ExitOnError)
	joinPort := joinFlag.String("port", defaultPort, "port number")
	bookID := joinFlag.String("bookid", "", "book id")

	serverFlag := flag.NewFlagSet("server", flag.ExitOnError)
	serverPort := serverFlag.String("port", defaultPort, "port number")
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
		command = &cmd.JoinCommand{
			OutputWriter:    outputWriter,
			PaperRepository: &client.PaperRepositoryImpl{Port: *joinPort},
			StateRepository: &client.StateRepositoryImpl{},
			BookID:          *bookID,
		}
	case "server":
		if err := serverFlag.Parse(args[1:]); err != nil {
			return nil, err
		}
		processor := server.NewProcessor()
		command = &cmd.ServerCommand{
			OutputWriter: outputWriter,
			Port:         *serverPort,
			ConfigPath:   *configPath,
			Processor:    processor,
			PaperController: controller.PaperController{
				PaperRepository: &server.PaperRepositoryImpl{
					Processor: processor,
					BookRepository: &server.BookRepositoryImpl{
						Processor: processor,
					},
				},
			},
		}
	default:
		return nil, fmt.Errorf("Not found command: %v", args[0])
	}

	return command, nil
}
