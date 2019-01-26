package datastore

// BookImpl implements book
type BookImpl struct {
	BookID string `json:"id"`
}

// ID returns an id
func (p *BookImpl) ID() string {
	return p.BookID
}
