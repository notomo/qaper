package datastore

// PaperImpl implements paper
type PaperImpl struct {
	PaperID string `json:"id"`
}

// ID returns an id
func (p *PaperImpl) ID() string {
	return p.PaperID
}
