package datastore

import "github.com/notomo/qaper/domain/model"

// PaperImpl implements paper
type PaperImpl struct {
	PaperID   string     `json:"id"`
	PaperBook model.Book `json:"book"`
}

// ID returns an id
func (p *PaperImpl) ID() string {
	return p.PaperID
}

// Book returns an book
func (p *PaperImpl) Book() model.Book {
	return p.PaperBook
}
