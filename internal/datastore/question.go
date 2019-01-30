package datastore

// QuestionImpl implements question
type QuestionImpl struct {
	QuestionID   string `json:"id"`
	QuestionBody string `json:"body"`
}

// ID returns an id
func (p *QuestionImpl) ID() string {
	return p.QuestionID
}

// Body returns an body
func (p *QuestionImpl) Body() string {
	return p.QuestionBody
}
