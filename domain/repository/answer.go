package repository

import "github.com/notomo/qaper/domain/model"

// AnswerRepository provides answer operations.
type AnswerRepository interface {
	Set(paperID string, answer model.Answer) error
}
