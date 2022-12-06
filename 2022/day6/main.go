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

	answer, err := search(string(input), 4)
	if err != nil {
		return fmt.Errorf("search error: %w", err)
	}

	fmt.Fprintln(stdout, "First 4-character marker ends at character", answer)

	answer, err = search(string(input), 14)
	if err != nil {
		return fmt.Errorf("search error: %w", err)
	}

	fmt.Fprintln(stdout, "First 14-character marker ends at character", answer)

	return nil
}

func search(s string, markerLength int) (int, error) {
	for i := markerLength - 1; i < len(s); i++ {
		m := make(map[uint8]struct{})
		for j := 0; j < markerLength; j++ {
			m[s[i-j]] = struct{}{}
		}
		if len(m) == markerLength {
			return i + 1, nil
		}
	}

	return 0, errors.New("first character not found")
}
