package datastore

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/internal"
	"github.com/notomo/qaper/internal/datastore"
)

// Processor represents a server
type Processor struct {
	addedPaper chan *datastore.PaperImpl
	papers     map[string]*datastore.PaperImpl

	gotPaper chan string
	paper    chan *datastore.PaperImpl

	gotBook chan string
	book    chan *datastore.BookImpl
	books   map[string]*datastore.BookImpl

	setAnswer   chan error
	paperAnswer chan *paperAnswer

	done chan bool
}

// NewProcessor create a processor
func NewProcessor() *Processor {
	addedPaper := make(chan *datastore.PaperImpl)
	papers := make(map[string]*datastore.PaperImpl)

	gotPaper := make(chan string)
	paper := make(chan *datastore.PaperImpl)

	gotBook := make(chan string)
	book := make(chan *datastore.BookImpl)
	books := make(map[string]*datastore.BookImpl)

	setAnswer := make(chan error)
	paperAnswer := make(chan *paperAnswer)

	done := make(chan bool)

	return &Processor{
		addedPaper: addedPaper,
		papers:     papers,

		gotPaper: gotPaper,
		paper:    paper,

		gotBook: gotBook,
		book:    book,
		books:   books,

		setAnswer:   setAnswer,
		paperAnswer: paperAnswer,

		done: done,
	}
}

// LoadLibrary loads the library books
func (processor *Processor) LoadLibrary(libraryPath string) error {
	if libraryPath == "" {
		return nil
	}

	abspath, err := filepath.Abs(libraryPath)
	if err != nil {
		return err
	}

	r, err := ioutil.ReadFile(abspath)
	if err != nil {
		return err
	}

	var library map[string]*datastore.BookImpl
	if err := json.Unmarshal(r, &library); err != nil {
		return err
	}
	processor.books = library
	return nil
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
			book, _ := processor.books[bookID]
			processor.book <- book
		case paperID := <-processor.gotPaper:
			paper, _ := processor.papers[paperID]
			processor.paper <- paper
		case paperAnswer := <-processor.paperAnswer:
			err := paperAnswer.paper.SetAnswer(paperAnswer.answer)
			processor.setAnswer <- err
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

// GetBook emmits a get book event
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

// GetPaper emmits a get paper event
func (processor *Processor) GetPaper(paperID string) (model.Paper, error) {
	go func() {
		processor.gotPaper <- paperID
	}()

	paper := <-processor.paper
	if paper == nil {
		return nil, internal.ErrNotFound
	}
	return paper, nil
}

type paperAnswer struct {
	paper  model.Paper
	answer model.Answer
}

// SetAnswer emmits a set answer event
func (processor *Processor) SetAnswer(paperID string, answer model.Answer) error {
	paper, err := processor.GetPaper(paperID)
	if err != nil {
		return err
	}

	go func() {
		processor.paperAnswer <- &paperAnswer{paper, answer}
	}()

	return <-processor.setAnswer
}
