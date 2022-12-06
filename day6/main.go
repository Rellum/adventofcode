package main

import (
	"errors"
	"fmt"
	"io"
	"os"
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
	input, err := io.ReadAll(stdin)
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	answer, err := part1(string(input))
	if err != nil {
		return fmt.Errorf("search error: %w", err)
	}

	fmt.Fprintln(stdout, "First marker after character", answer)

	return nil
}

func part1(s string) (int, error) {
	for i := 3; i < len(s); i++ {
		m := make(map[uint8]struct{})
		for j := 0; j < 4; j++ {
			m[s[i-j]] = struct{}{}
		}
		if len(m) == 4 {
			return i + 1, nil
		}
	}

	return 0, errors.New("first character not found")
}
