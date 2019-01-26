package repository

import (
	"github.com/notomo/qaper/domain/model"
)

// BookRepository provides book operations.
type BookRepository interface {
	Get(bookID string) (model.Book, error)
}
