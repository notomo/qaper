package repository

import (
	"github.com/notomo/qaper/domain/model"
)

// PaperRepository provides paper operations.
type PaperRepository interface {
	Add() (model.Paper, error)
}
