package datastore

import (
	"github.com/notomo/qaper/domain/model"
)

// BookRepositoryImpl implements book repository
type BookRepositoryImpl struct {
	Processor *Processor
}

// Get a book
func (repo *BookRepositoryImpl) Get(bookID string) (model.Book, error) {
	return repo.Processor.GetBook(bookID)
}
