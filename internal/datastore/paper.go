package datastore

import (
	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/internal"
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
	if len(p.Answers) <= p.PaperCurrentIndex {
		return internal.ErrOutOfRange
	}
	answerImpl, _ := answer.(*AnswerImpl)
	p.Answers[p.PaperCurrentIndex] = answerImpl
	return nil
}
