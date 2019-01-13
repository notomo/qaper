package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
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
	id string
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

// JoinResponse represents a response on join
type JoinResponse struct {
	ID string `json:"id"`
}

// Add the client to the server
func (s *Server) Add(client *Client) {
	s.joined <- client
}

// Listen starts listening a port.
func (s *Server) Listen() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/paper", func(w http.ResponseWriter, r *http.Request) {
		client := &Client{id: "0"}
		s.Add(client)

		response := &JoinResponse{ID: "0"}
		res, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

	server := &http.Server{Addr: ":" + s.Port, Handler: mux}
	go func() {
		server.ListenAndServe()
	}()

	for {
		select {
		case client := <-s.joined:
			if _, ok := s.clients[client.id]; !ok {
				s.clients[client.id] = client
				log.Printf("Joined: %v\n", client.id)
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
