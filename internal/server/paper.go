package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/notomo/qaper/internal/datastore"
)

func (s *Server) paper(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	client := &Client{ID: "0"}
	s.Add(client)

	paper := datastore.PaperImpl{PaperID: client.ID}

	responseJSON(w, paper)
}
