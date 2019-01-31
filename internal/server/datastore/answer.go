package datastore

import (
	"github.com/notomo/qaper/domain/model"
)

// AnswerRepositoryImpl implements answer repository
type AnswerRepositoryImpl struct {
	Processor *Processor
}

// Set sets an answer
func (repo *AnswerRepositoryImpl) Set(paperID string, answer model.Answer) error {
	return repo.Processor.SetAnswer(paperID, answer)
}
