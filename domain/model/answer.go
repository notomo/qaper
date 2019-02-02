package model

import "io"

// Answer represents a question answer
type Answer interface {
	Body() string
}

// AnswerDecoder decodes answers
type AnswerDecoder interface {
	Decode(io.Reader) (Answer, error)
}
