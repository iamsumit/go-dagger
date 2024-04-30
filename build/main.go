// A generated module for GoDagger functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import "context"

const (
	// source build directory
	workdirSrc = "/src"

	// used for applicaiton runtime
	workdirRuntime = "/app"

	// build file
	buildFile = "build"
)

type GoDagger struct {
	Greeting string

	// Src is the source directory where the main go.mod is located.
	Src *Directory

	// BuildCtr is the container to use for the build.
	BuildCtr *Container

	// RunCtr is the container to use for the run.
	RunCtr *Container
}

// New creates a new instance of GoDagger
func New(src *Directory) *GoDagger {
	buildCtr := dag.Container().
		From("golang:1.21-alpine").
		WithWorkdir(workdirSrc).
		WithDirectory(workdirSrc, src)

	runCtr := dag.Container().
		From("golang:1.21-alpine").
		WithWorkdir(workdirRuntime)

	return &GoDagger{
		Greeting: "hello",
		Src:      src,
		BuildCtr: buildCtr,
		RunCtr:   runCtr,
	}
}

// Output prints out the built staging output
func (m *GoDagger) Output(
	ctx context.Context,
) (string, error) {
	return m.BuildCtr.Stdout(ctx)
}

// Deps prepares the module for build by downloading dependencies
func (m *GoDagger) Deps() *GoDagger {
	m.BuildCtr = m.BuildCtr.
		WithExec([]string{"go", "mod", "download"})

	return m
}

// Test runs the tests for the module
func (m *GoDagger) Test() *GoDagger {
	m.BuildCtr = m.BuildCtr.
		WithDirectory(workdirSrc, m.Src).
		WithExec([]string{"go", "test", "./..."})

	return m
}

// Build builds the module
func (m *GoDagger) Build() *GoDagger {
	m.BuildCtr = m.BuildCtr.
		WithExec([]string{"go", "build", "-o", buildFile})

	return m
}

// Publish the build image
func (m *GoDagger) Publish(ctx context.Context) (string, error) {
	return m.RunCtr.
		WithFiles(workdirRuntime, []*File{
			m.BuildCtr.File(buildFile),
		}).
		Publish(ctx, "ttl.sh/go-dagger-container")
}
