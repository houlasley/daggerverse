// A generated module for Cdk functions
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

	"dagger/cdk/internal/dagger"
)

type CDK struct{}

// Returns a container that echoes whatever string argument is provided
func (m *CDK) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *CDK) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

func (m *CDK) Synth(ctx context.Context, appDir *dagger.Directory) (*dagger.Directory, error) {
	return dag.Container().
		From("python:3.11").
		WithExec([]string{"pip", "install", "uv"}).
		WithMountedDirectory("/app", appDir).
		WithWorkdir("/app").
		WithExec([]string{"uv", "pip", "install"}).
		WithExec([]string{"npm", "install", "-g", "aws-cdk"}).
		WithExec([]string{"cdk", "synth"}).
		Directory("cdk.out"), nil
}

func (m *CDK) Deploy(ctx context.Context, appDir *dagger.Directory, env, accountID string) error {
	_, err := dag.Container().
		From("python:3.11").
		WithExec([]string{"pip", "install", "uv"}).
		WithMountedDirectory("/app", appDir).
		WithWorkdir("/app").
		WithEnvVariable("AWS_REGION", "us-east-1").
		WithEnvVariable("AWS_ACCOUNT", accountID).
		WithExec([]string{"uv", "pip", "install"}).
		WithExec([]string{"npm", "install", "-g", "aws-cdk"}).
		WithExec([]string{"cdk", "deploy", "--require-approval", "never"}).
		Sync(ctx)
	return err
}

func (m *CDK) Destroy(ctx context.Context, appDir *dagger.Directory, env, accountID string) error {
	_, err := dag.Container().
		From("python:3.11").
		WithExec([]string{"pip", "install", "uv"}).
		WithMountedDirectory("/app", appDir).
		WithWorkdir("/app").
		WithEnvVariable("AWS_REGION", "us-east-1").
		WithEnvVariable("AWS_ACCOUNT", accountID).
		WithExec([]string{"uv", "pip", "install"}).
		WithExec([]string{"npm", "install", "-g", "aws-cdk"}).
		WithExec([]string{"cdk", "destroy", "--force"}).
		Sync(ctx)
	return err
}
