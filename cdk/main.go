package main

import (
	"context"
	"fmt"

	"dagger/cdk/internal/dagger"
)

type Cdk struct{}

// `Main()` is the entry point for the Dagger execution
func (m *Cdk) Build(ctx context.Context) (*dagger.Container, error) {
	// 🔥 Build a Wolfi-based image with both Python & Node.js
	finalImage := BuildWolfiImage(ctx, dag)

	// 🛠 Verify Python & Node.js installation

	// 🛠 Verify Python & Node.js installation

	return finalImage, nil
}

// 🔥 Build a minimal Wolfi-based image with both Python & Node.js
func BuildWolfiImage(ctx context.Context, dag *dagger.Client) *dagger.Container {
	// 🟢 Stage 1: Pull Wolfi's Node.js image
	dag.Container().
		From("cgr.dev/chainguard/node:latest").
		Terminal()

	// 🟢 Stage 2: Pull Wolfi's Python image
	python := dag.Container().
		From("cgr.dev/chainguard/python:latest").
		Terminal().
		WithExec([]string{"python", "--version"}) // Ensure Python is available

	fmt.Println(python)
	fmt.Println(python)
	fmt.Println(python)
	fmt.Println(python)
	fmt.Println(python)
	fmt.Println(python)
	fmt.Println(python)
	fmt.Println(python)

	// 🏗 Final Stage: Combine both into a minimal Wolfi-based image
	finalImage := dag.Container().
		From("cgr.dev/chainguard/wolfi-base:latest").
		// Copy Node.js files explicitly
		WithExec([]string{"ls"})

	return finalImage
}
