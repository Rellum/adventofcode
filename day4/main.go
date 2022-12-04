package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

type assignment struct {
	from, to int
}

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	var fullyContainedCount int

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return fmt.Errorf("unexpected empty line")
		}

		assignments := strings.SplitN(line, ",", 2)

		assignment1, err := parseAssignment(assignments[0])
		if err != nil {
			return fmt.Errorf("failed to parse first assigment '%s': %w", assignments[0], err)
		}

		assignment2, err := parseAssignment(assignments[1])
		if err != nil {
			return fmt.Errorf("failed to parse second assigment '%s': %w", assignments[1], err)
		}

		if assignment1.contains(assignment2) || assignment2.contains(assignment1) {
			fullyContainedCount++
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Fprintln(stdout, "Count:", fullyContainedCount)

	return nil
}

func parseAssignment(s string) (*assignment, error) {
	bounds := strings.SplitN(s, "-", 2)

	from, err := strconv.Atoi(bounds[0])
	if err != nil {
		return nil, fmt.Errorf("parsing the lower bound '%s': %w", bounds[0], err)
	}

	to, err := strconv.Atoi(bounds[1])
	if err != nil {
		return nil, fmt.Errorf("parsing the higher bound '%s': %w", bounds[1], err)
	}

	return &assignment{from, to}, nil
}

func (a *assignment) contains(b *assignment) bool {
	if a.from > b.from {
		return false
	}

	if a.to < b.to {
		return false
	}

	return true
}
