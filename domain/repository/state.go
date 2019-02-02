package repository

import (
	"github.com/notomo/qaper/domain/model"
)

// StateRepository provides state operations.
type StateRepository interface {
	Save(paper model.Paper) error
	Load() (model.State, error)
}
