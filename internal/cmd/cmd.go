package cmd

// Command represents a command
type Command interface {
	Run() error
}
