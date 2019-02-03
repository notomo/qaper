package repository

import (
	"github.com/notomo/qaper/domain/model"
	cmodel "github.com/notomo/qaper/internal/client/model"
)

// StateRepository provides state operations.
type StateRepository interface {
	Save(paper model.Paper) error
	Load() (cmodel.State, error)
}
