package model

// Book represents a question book
type Book interface {
	ID() string
	Questions() []Question
}
