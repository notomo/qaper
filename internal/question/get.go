package question

// Question represents a question
type Question struct {
	Name string
}

func (q *Question) String() string {
	return q.Name
}

// Get a question
func Get() (*Question, error) {
	return &Question{Name: "test"}, nil
}
