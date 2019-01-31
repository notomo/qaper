package cmd

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/notomo/qaper/internal/server/api/controller"
	"github.com/notomo/qaper/internal/server/datastore"
)

// ServerCommand represents `server` command
type ServerCommand struct {
	OutputWriter    io.Writer
	Port            string
	LibraryPath     string
	Processor       *datastore.Processor
	PaperController controller.PaperController
}

// Run `server` command
func (c *ServerCommand) Run() error {
	router := httprouter.New()
	router.POST("/book/:bookID/paper", c.PaperController.Add)
	router.GET("/paper/:paperID/question", c.PaperController.GetCurrentQuestion)
	server := &http.Server{
		Addr:    ":" + c.Port,
		Handler: router,
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	if err := c.Processor.LoadLibrary(c.LibraryPath); err != nil {
		return err
	}
	go c.Processor.Start()

	log.Println("Listen: " + c.Port)
	return server.ListenAndServe()
}
