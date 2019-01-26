package datastore

import (
	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/internal/datastore"
)

// PaperRepositoryImpl implements paper repository
type PaperRepositoryImpl struct {
	Processor *Processor
}

// Add adds a paper
func (repo *PaperRepositoryImpl) Add() (model.Paper, error) {
	paper := &datastore.PaperImpl{PaperID: "0"}
	repo.Processor.AddPaper(paper)
	return paper, nil
}
