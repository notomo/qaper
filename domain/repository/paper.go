package repository

import (
	"github.com/notomo/qaper/domain/model"
)

// PaperRepository provides paper operations.
type PaperRepository interface {
	Add(bookID string) (model.Paper, error)
	Get(paperID string) (model.Paper, error)
}
