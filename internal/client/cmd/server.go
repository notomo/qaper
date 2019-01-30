package cmd

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/notomo/qaper/internal/client/cmd/config"
	"github.com/notomo/qaper/internal/server/api/controller"
	"github.com/notomo/qaper/internal/server/datastore"
)

// ServerCommand represents `server` command
type ServerCommand struct {
	OutputWriter    io.Writer
	Port            string
	LibraryPath     string
	ConfigPath      string
	Processor       *datastore.Processor
	PaperController controller.PaperController
}

// Run `server` command
func (c *ServerCommand) Run() error {
	serverConfig, err := (&config.ServerConfig{
		LibraryPath: c.LibraryPath,
		Port:        c.Port,
	}).Load(c.ConfigPath)
	if err != nil {
		return err
	}

	router := httprouter.New()
	router.POST("/book/:bookID/paper", c.PaperController.Add)
	router.GET("/paper/:paperID/question", c.PaperController.GetCurrentQuestion)
	server := &http.Server{
		Addr:    ":" + serverConfig.Port,
		Handler: router,
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	if err := c.Processor.LoadLibrary(serverConfig.LibraryPath); err != nil {
		return err
	}
	go c.Processor.Start()

	log.Println("Listen: " + serverConfig.Port)
	return server.ListenAndServe()
}
