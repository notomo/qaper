package model

// Paper represents a question paper
type Paper interface {
	ID() string
	CurrentQuestion() Question
	SetAnswer(answer Answer) error
}
