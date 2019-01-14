package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Server represents a server
type Server struct {
	Port    string
	done    chan bool
	joined  chan *Client
	clients map[string]*Client
}

// Client represents a client
type Client struct {
	ID string
}

// Config represents a server configuration
type Config struct {
	Port   string
	server *Server
}

// Server returns a server
func (c *Config) Server() *Server {
	if c.server != nil {
		return c.server
	}

	done := make(chan bool)
	joined := make(chan *Client)
	clients := make(map[string]*Client)
	return &Server{Port: c.Port, done: done, joined: joined, clients: clients}
}

// Add the client to the server
func (s *Server) Add(client *Client) {
	s.joined <- client
}

func responseJSON(w http.ResponseWriter, v interface{}) {
	res, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// Listen starts listening a port.
func (s *Server) Listen() error {
	router := httprouter.New()
	router.POST("/paper", s.paper)

	server := &http.Server{Addr: ":" + s.Port, Handler: router}
	go func() {
		log.Println("Listen: " + s.Port)
		server.ListenAndServe()
	}()

	for {
		select {
		case client := <-s.joined:
			if _, ok := s.clients[client.ID]; !ok {
				s.clients[client.ID] = client
				log.Printf("Joined: %v\n", client.ID)
			}
		case <-s.done:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			server.Shutdown(ctx)
			log.Println("Done")
			return nil
		}
	}
}
