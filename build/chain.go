package main

import "context"

// Update the default greet string
func (m *GoDagger) NewGreet(ctx context.Context) *GoDagger {
	m.Greeting = "hey hello"
	return m
}

// print the greeting string
func (m *GoDagger) PrintGreet(ctx context.Context) (string, error) {
	return m.Greeting, nil
}
