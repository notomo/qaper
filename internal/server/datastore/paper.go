package datastore

import (
	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/domain/repository"
	"github.com/notomo/qaper/internal/datastore"
	"github.com/rs/xid"
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

	bookImpl, _ := book.(*datastore.BookImpl)
	paperID := xid.New().String()
	paper := &datastore.PaperImpl{PaperID: paperID, PaperBook: bookImpl}
	repo.Processor.AddPaper(paper)
	return paper, nil
}
