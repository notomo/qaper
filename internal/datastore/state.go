package datastore

// StateImpl implements state
type StateImpl struct {
	StatePaperID string `json:"paperId"`
}

// PaperID returns a paper id
func (p *StateImpl) PaperID() string {
	return p.StatePaperID
}
