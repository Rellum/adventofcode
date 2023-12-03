package main

import (
	"bufio"
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
	// Read all
	input, err := io.ReadAll(stdin)
	if err != nil {
		return fmt.Errorf("read all: %w", err)
	}
	_ = input

	// Read runes
	r := bufio.NewReader(stdin)
	for {
		c, size, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return fmt.Errorf("read rune: %w", err)
			}
		}

		_, _ = c, size

		// do stuff with a rune
	}

	// Scan lines
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var i int
		_, err := fmt.Sscanf(line, "Verbose puzzle input line with value y=%d in the middle of it", &i)
		if err != nil {
			return fmt.Errorf("error parsing line '%s': %w", line, err)
		}

		// do stuff with a line
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Fprintf(stdout, "Answer: %d\n", 123)

	return nil
}
