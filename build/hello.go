package main

import "context"

// Print a message to the console
func (m *GoDagger) Echo(stringArg string) *Container {
	return dag.
		Container().
		From("alpine:latest").
		WithExec([]string{"echo", stringArg})
}

// Message is a struct that contains a message and its source
type Message struct {
	// The message to be printed
	Text string

	// The source of the message
	From string
}

// Returns a Message struct with a default text and source
func (m *GoDagger) Message() Message {
	return Message{
		Text: "Hello, world!",
		From: "Dagger",
	}
}

// Print hello world.
func (m *GoDagger) Hello(
	ctx context.Context,
	// Change the greeting
	// +optional
	// +default="hello"
	greeting string,
) (string, error) {
	return dag.
		Hello().
		Hello(ctx, HelloHelloOpts{
			Greeting: greeting,
		})
}
