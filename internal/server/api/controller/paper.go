package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/notomo/qaper/domain/repository"
	"github.com/notomo/qaper/internal"
	"github.com/notomo/qaper/internal/datastore"
)

// PaperController responds to paper api
type PaperController struct {
	PaperRepository  repository.PaperRepository
	AnswerRepository repository.AnswerRepository
}

// Add adds a paper
func (ctrl *PaperController) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bookID := params.ByName("bookID")
	paper, err := ctrl.PaperRepository.Add(bookID)
	if err == internal.ErrNotFound {
		response404(w, fmt.Sprintf("Not Found book: %v", bookID))
		return
	}
	if err != nil {
		response500(w, err.Error())
		return
	}

	responseJSON(w, paper)
}

// Get gets a paper
func (ctrl *PaperController) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	paperID := params.ByName("paperID")
	paper, err := ctrl.PaperRepository.Get(paperID)
	if err == internal.ErrNotFound {
		response404(w, fmt.Sprintf("Not Found paper: %v", paperID))
		return
	}
	if err != nil {
		response500(w, err.Error())
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

// SetAnswer gets a current question
func (ctrl *PaperController) SetAnswer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	// TODO: answer decoder interface
	var answer datastore.AnswerImpl
	if err := decoder.Decode(&answer); err != nil {
		response400(w, fmt.Sprintf("json parse error"))
		return
	}

	paperID := params.ByName("paperID")
	err := ctrl.AnswerRepository.Set(paperID, &answer)
	if err == internal.ErrNotFound {
		response404(w, fmt.Sprintf("Not Found paper: %v", paperID))
		return
	}
	if err != nil {
		response500(w, err.Error())
		return
	}

	responseJSON(w, answer)
}
