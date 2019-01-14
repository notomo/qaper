package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) paper(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	client := &Client{ID: "0"}
	s.Add(client)

	w.Write([]byte(client.ID))
}
