package server

import (
	"net/http"
)

func (s *Server) paper(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	client := &Client{ID: "0"}
	s.Add(client)

	w.Write([]byte(client.ID))
}
