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

import (
	"context"
)

type GoDagger struct{}

// Returns a container that echoes whatever string argument is provided
func (m *GoDagger) ContainerEcho(stringArg string) *Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *GoDagger) GrepDir(ctx context.Context, directoryArg *Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
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

// Test go app.
func (m *GoDagger) Test(
	ctx context.Context,
	dir *Directory,
) (string, error) {
	return dag.
		Golang().
		Test(ctx, GolangTestOpts{
			Source: dir,
			Args:   []string{"./..."},
		})
}

// Build and publish a project using a Wolfi container
func (m *GoDagger) BuildAndPublish(ctx context.Context, dir *Directory, args []string) (string, error) {
	ctr := dag.Wolfi().Container()

	return dag.
		Golang().
		BuildContainer(GolangBuildContainerOpts{Source: dir, Args: args, Base: ctr}).
		Publish(ctx, "ttl.sh/my-hello-container") //#nosec
}
