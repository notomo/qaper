package datastore

import (
	"github.com/notomo/qaper/domain/model"
)

// BookImpl implements book
type BookImpl struct {
	BookID        string          `json:"id"`
	BookQuestions []*QuestionImpl `json:"questions"`
}

// ID returns an id
func (p *BookImpl) ID() string {
	return p.BookID
}

// Questions returns questions
func (p *BookImpl) Questions() []model.Question {
	questions := make([]model.Question, len(p.BookQuestions))
	for i, question := range p.BookQuestions {
		questions[i] = question
	}
	return questions
}
