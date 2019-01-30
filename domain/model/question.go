package model

// Question represents a question
type Question interface {
	ID() string
	Body() string
}
