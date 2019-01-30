package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/notomo/qaper/domain/repository"
	"github.com/notomo/qaper/internal"
)

// PaperController responds to paper api
type PaperController struct {
	PaperRepository repository.PaperRepository
}

// Add adds a paper
func (ctrl *PaperController) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bookID := params.ByName("bookID")
	paper, err := ctrl.PaperRepository.Add(bookID)
	if err == internal.ErrNotFound {
		response404(w, fmt.Sprintf("Not Found book: %v", bookID))
		return
	}

	responseJSON(w, paper)
}

// GetCurrentQuestion gets a current question
func (ctrl *PaperController) GetCurrentQuestion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	paperID := params.ByName("paperID")
	paper, err := ctrl.PaperRepository.Get(paperID)
	if err == internal.ErrNotFound {
		response404(w, fmt.Sprintf("Not Found paper: %v", paperID))
		return
	}

	question := paper.CurrentQuestion()
	if question == nil {
		response404(w, fmt.Sprintf("Not Found question on paper: %v", paperID))
		return
	}

	responseJSON(w, question)
}
