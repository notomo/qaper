package datastore

import (
	"log"

	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/internal"
	"github.com/notomo/qaper/internal/datastore"
)

// Processor represents a server
type Processor struct {
	addedPaper chan *datastore.PaperImpl
	papers     map[string]*datastore.PaperImpl

	gotBook chan string
	book    chan *datastore.BookImpl
	books   map[string]*datastore.BookImpl

	done chan bool
}

// NewProcessor create a processor
func NewProcessor() *Processor {
	addedPaper := make(chan *datastore.PaperImpl)
	papers := make(map[string]*datastore.PaperImpl)

	gotBook := make(chan string)
	book := make(chan *datastore.BookImpl)
	books := make(map[string]*datastore.BookImpl)

	done := make(chan bool)

	return &Processor{
		addedPaper: addedPaper,
		papers:     papers,

		gotBook: gotBook,
		book:    book,
		books:   books,

		done: done,
	}
}

// Start process
func (processor *Processor) Start() error {
	for {
		select {
		case paper := <-processor.addedPaper:
			if _, ok := processor.papers[paper.ID()]; !ok {
				processor.papers[paper.ID()] = paper
				log.Printf("addedPaper: %v\n", paper.ID())
			}
		case bookID := <-processor.gotBook:
			book, ok := processor.books[bookID]
			if !ok {
				processor.book <- book
			} else {
				processor.book <- nil
			}
		case <-processor.done:
			log.Println("Done")
			return nil
		}
	}
}

// AddPaper emmits an add paper event
func (processor *Processor) AddPaper(paper *datastore.PaperImpl) {
	processor.addedPaper <- paper
}

// GetBook emmits an add paper event
func (processor *Processor) GetBook(bookID string) (model.Book, error) {
	go func() {
		processor.gotBook <- bookID
	}()

	book := <-processor.book
	if book == nil {
		return nil, internal.ErrNotFound
	}
	return book, nil
}
