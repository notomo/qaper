package datastore

// AnswerImpl implements answer
type AnswerImpl struct {
	AnswerBody string `json:"body"`
}

// Body returns an answer body
func (p *AnswerImpl) Body() string {
	return p.AnswerBody
}
