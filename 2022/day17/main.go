package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	jetPattern, err := io.ReadAll(stdin)
	if err != nil {
		return fmt.Errorf("read all: %w", err)
	}

	c := cave{rows: []string{
		strings.Repeat(".", 7),
		strings.Repeat(".", 7),
		strings.Repeat(".", 7),
		strings.Repeat(".", 7),
	}}

	fmt.Fprintf(stdout, "Answer: %d\n", 123)

	return nil
}

type cave struct {
	rows []string
}
