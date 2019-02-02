package internal

import "errors"

var (
	// ErrNotFound represents a not found error
	ErrNotFound = errors.New("Not Found")
	// ErrNotImplemented represents the function is not implemented
	ErrNotImplemented = errors.New("Not Implemented")
	// ErrOutOfRange represents the index is out of range.
	ErrOutOfRange = errors.New("Out of range")
)
