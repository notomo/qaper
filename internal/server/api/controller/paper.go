package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/domain/repository"
	"github.com/notomo/qaper/internal"
)

// PaperController responds to paper api
type PaperController struct {
	PaperRepository  repository.PaperRepository
	AnswerRepository repository.AnswerRepository
	AnswerDecoder    model.AnswerDecoder
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
	answer, err := ctrl.AnswerDecoder.Decode(r.Body)
	if err != nil {
		response400(w, err.Error())
		return
	}

	paperID := params.ByName("paperID")
	if err := ctrl.AnswerRepository.Set(paperID, answer); err != nil {
		switch err {
		case internal.ErrNotFound:
			response404(w, fmt.Sprintf("Not Found paper: %v", paperID))
		default:
			response500(w, err.Error())
		}
		return
	}

	responseJSON(w, answer)
}
