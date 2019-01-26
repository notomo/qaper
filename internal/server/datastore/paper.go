package datastore

import (
	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/domain/repository"
	"github.com/notomo/qaper/internal/datastore"
)

// PaperRepositoryImpl implements paper repository
type PaperRepositoryImpl struct {
	Processor      *Processor
	BookRepository repository.BookRepository
}

// Add adds a paper
func (repo *PaperRepositoryImpl) Add(bookID string) (model.Paper, error) {
	book, err := repo.BookRepository.Get(bookID)
	if err != nil {
		return nil, err
	}

	paper := &datastore.PaperImpl{PaperID: "0", PaperBook: book}
	repo.Processor.AddPaper(paper)
	return paper, nil
}
