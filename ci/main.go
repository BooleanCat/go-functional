package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func build(ctx context.Context) error {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return err
	}
	defer client.Close()

	src := client.Host().Directory(".")

	_, err = client.Container().From(
		"golang:1.22",
	).WithDirectory(
		"/src", src,
	).WithWorkdir(
		"/src",
	).WithEnvVariable(
		"GOEXPERIMENT", "rangefunc",
	).WithEnvVariable(
		"SKIP_LINT", "true",
	).WithExec(
		[]string{"make", "check"},
	).Stdout(ctx)

	return err
}
