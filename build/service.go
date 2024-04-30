package main

import "context"

// HttpService returns a service that runs a simple HTTP server
func (m *GoDagger) HttpService() *Service {
	return dag.
		Container().
		From("python:latest").
		WithWorkdir("/srv").
		WithNewFile("index.html", ContainerWithNewFileOpts{
			Contents: "Hello, world!",
		}).
		WithExec([]string{"python", "-m", "http.server", "8080"}).
		WithExposedPort(8080).
		AsService()
}

// sends a request to an HTTP service and returns the response
func (m *GoDagger) Get(ctx context.Context) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithServiceBinding("www", m.HttpService()).
		WithExec([]string{"wget", "-O-", "http://www:8080"}).
		Stdout(ctx)
}
