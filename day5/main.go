package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
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

var movePattern = regexp.MustCompile(`^move ([0-9]+) from ([0-9]+) to ([0-9]+)$`)

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	var stacks [][]rune
	scanner := bufio.NewScanner(stdin)

	// Parse starting position
	stacks, err := parseStacks(scanner)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	// Empty line
	scanner.Scan()
	if len(scanner.Text()) != 0 {
		return fmt.Errorf("expected empty line after the stack starting state")
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	// Moves
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return fmt.Errorf("unexpected empty line")
		}

		matches := movePattern.FindStringSubmatch(line)

		if len(matches) != 4 {
			return fmt.Errorf("move instruction could not be parsed: '%s'", line)
		}

		count, err := strconv.Atoi(matches[1])
		if err != nil {
			return fmt.Errorf("error parsing '%s' as a number", matches[1])
		}

		from, err := strconv.Atoi(matches[2])
		if err != nil {
			return fmt.Errorf("error parsing '%s' as a number", matches[2])
		}
		from-- // we use zero-indexed stack numbering

		to, err := strconv.Atoi(matches[3])
		if err != nil {
			return fmt.Errorf("error parsing '%s' as a number", matches[3])
		}
		to-- // we use zero-indexed stack numbering

		for i := 0; i < count; i++ {
			var popped rune
			popped, stacks[from] = stacks[from][0], stacks[from][1:]
			stacks[to] = append([]rune{popped}, stacks[to]...)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	var answer string
	for i := range stacks {
		answer += string(stacks[i][0])
	}

	fmt.Fprintln(stdout, "Top crates:", answer)

	return nil
}

func parseStacks(scanner *bufio.Scanner) ([][]rune, error) {
	var stacks [][]rune

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return nil, fmt.Errorf("unexpected empty line")
		}

		var capture bool
		var stack int
		for i, s := range line {
			if s == '1' {
				return stacks, scanner.Err()
			}

			if s == '[' {
				capture = true
				stack = i / 4
				continue
			}

			if s == ']' {
				capture = false
			}

			if !capture {
				continue
			}

			for len(stacks) <= stack {
				stacks = append(stacks, []rune{})
			}

			stacks[stack] = append(stacks[stack], s)
		}
	}

	return nil, errors.New("did not completely parse stack starting state")
}
