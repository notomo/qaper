package datastore

import (
	"encoding/json"
	"io"

	"github.com/notomo/qaper/domain/model"
)

// AnswerImpl implements answer
type AnswerImpl struct {
	AnswerBody string `json:"body"`
}

// Body returns an answer body
func (p *AnswerImpl) Body() string {
	return p.AnswerBody
}

// AnswerJSONDecoder implements answer decoder
type AnswerJSONDecoder struct{}

// Decode json answer
func (*AnswerJSONDecoder) Decode(reader io.Reader) (model.Answer, error) {
	decoder := json.NewDecoder(reader)

	var answer AnswerImpl
	if err := decoder.Decode(&answer); err != nil {
		return nil, err
	}

	return &answer, nil
}
