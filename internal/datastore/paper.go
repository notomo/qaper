package datastore

import (
	"github.com/notomo/qaper/domain/model"
)

// PaperImpl implements paper
type PaperImpl struct {
	PaperID           string    `json:"id"`
	PaperBook         *BookImpl `json:"book"`
	PaperCurrentIndex int       `json:"currentIndex"`
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
