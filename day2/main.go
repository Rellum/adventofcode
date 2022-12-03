package main

import (
	"bufio"
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

type shape int

const (
	Rock shape = iota + 1
	Paper
	Scissors
)

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	var part1Score int
	var part2Score int
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) != 3 {
			return fmt.Errorf("failed to parse '%s'. %d segments instead of 3", line, len(line))
		}

		if line[1] != ' ' {
			return fmt.Errorf("failed to parse '%s'. second character is not whitespace", line)
		}

		them, err := parseShape(line[0])
		if err != nil {
			return fmt.Errorf("failed to parse their move in '%s': %w", line, err)
		}

		me, err := parseShape(line[2])
		if err != nil {
			return fmt.Errorf("failed to parse my move in '%s': %w", line, err)
		}

		calculated, err := reply(them, line[2])
		if err != nil {
			return fmt.Errorf("failed to calculate reply in '%s': %w", line, err)
		}

		part1Score += int(me) + outcome(me, them)
		part2Score += int(calculated) + outcome(calculated, them)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Fprintln(stdout, "Score (part 1):", part1Score)
	fmt.Fprintln(stdout, "Score (part 2):", part2Score)

	return nil
}

func parseShape(c uint8) (shape, error) {
	switch c {
	case 'A', 'X':
		return Rock, nil
	case 'B', 'Y':
		return Paper, nil
	case 'C', 'Z':
		return Scissors, nil
	}
	return 0, errors.New("unrecognised shape symbol")
}

func reply(them shape, outcome uint8) (shape, error) {
	if outcome == 'Y' {
		return them, nil // draw
	}

	if outcome == 'X' { // lose
		if them == Rock {
			return Scissors, nil
		} else if them == Paper {
			return Rock, nil
		} else { // them == Scissors
			return Paper, nil
		}
	}

	if outcome == 'Z' { // win
		if them == Rock {
			return Paper, nil
		} else if them == Paper {
			return Scissors, nil
		} else { // them == Scissors
			return Rock, nil
		}
	}

	return 0, errors.New("unrecognised outcome")
}

func outcome(me, them shape) (score int) {
	if me == them {
		return 3
	}

	if me == Rock && them == Scissors {
		return 6
	}

	if me == Paper && them == Rock {
		return 6
	}

	if me == Scissors && them == Paper {
		return 6
	}

	return 0
}
