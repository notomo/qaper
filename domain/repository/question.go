package repository

import (
	"github.com/notomo/qaper/domain/model"
)

// QuestionRepository provides question operations.
type QuestionRepository interface {
	GetCurrent(paperID string) (model.Question, error)
}
