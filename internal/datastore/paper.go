package datastore

import (
	"github.com/notomo/qaper/domain/model"
)

// PaperImpl implements paper
type PaperImpl struct {
	PaperID           string        `json:"id"`
	PaperBook         *BookImpl     `json:"book"`
	PaperCurrentIndex int           `json:"currentIndex"`
	Answers           []*AnswerImpl `json:"answers"`
}

// ID returns an id
func (p *PaperImpl) ID() string {
	return p.PaperID
}

// CurrentQuestion returns a current question
func (p *PaperImpl) CurrentQuestion() model.Question {
	questions := p.PaperBook.Questions()
	if len(questions) <= p.PaperCurrentIndex {
		return nil
	}
	return questions[p.PaperCurrentIndex]
}

// SetAnswer sets an answer
func (p *PaperImpl) SetAnswer(answer model.Answer) error {
	answerImpl, _ := answer.(*AnswerImpl)
	p.Answers[p.PaperCurrentIndex] = answerImpl
	return nil
}
