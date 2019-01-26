package datastore

import (
	"log"

	"github.com/notomo/qaper/internal/datastore"
)

// Processor represents a server
type Processor struct {
	addedPaper chan *datastore.PaperImpl
	papers     map[string]*datastore.PaperImpl

	done chan bool
}

// NewProcessor create a processor
func NewProcessor() *Processor {
	addedPaper := make(chan *datastore.PaperImpl)
	papers := make(map[string]*datastore.PaperImpl)

	done := make(chan bool)

	return &Processor{
		addedPaper: addedPaper,
		papers:     papers,

		done: done,
	}
}

// AddPaper emmits an add paper event
func (processor *Processor) AddPaper(paper *datastore.PaperImpl) {
	processor.addedPaper <- paper
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
		case <-processor.done:
			log.Println("Done")
			return nil
		}
	}
}
