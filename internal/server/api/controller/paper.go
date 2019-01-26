package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/notomo/qaper/domain/repository"
)

// PaperController responds to /paper/...
type PaperController struct {
	PaperRepository repository.PaperRepository
}

// Add adds a paper
func (ctrl *PaperController) Add(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	paper, _ := ctrl.PaperRepository.Add()
	responseJSON(w, paper)
}
