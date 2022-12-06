package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
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

var movePattern = regexp.MustCompile(`^move ([0-9]+) from ([0-9]+) to ([0-9]+)$`)

type move struct {
	from, to, count int
}

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	var stacks [][]rune
	var moves []move
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

		var m move

		m.count, err = strconv.Atoi(matches[1])
		if err != nil {
			return fmt.Errorf("error parsing '%s' as a number", matches[1])
		}

		m.from, err = strconv.Atoi(matches[2])
		if err != nil {
			return fmt.Errorf("error parsing '%s' as a number", matches[2])
		}
		m.from-- // we use zero-indexed stack numbering

		m.to, err = strconv.Atoi(matches[3])
		if err != nil {
			return fmt.Errorf("error parsing '%s' as a number", matches[3])
		}
		m.to-- // we use zero-indexed stack numbering

		moves = append(moves, m)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	// Part 1 calculation
	part1stacks := cloneStacks(stacks)

	for i := range moves {
		m := moves[i]
		for i := 0; i < m.count; i++ {
			var popped rune
			popped, part1stacks[m.from] = part1stacks[m.from][0], part1stacks[m.from][1:]
			part1stacks[m.to] = append([]rune{popped}, part1stacks[m.to]...)
		}
	}

	fmt.Fprint(stdout, "Top crates (part 1): ")
	for i := range stacks {
		fmt.Fprint(stdout, string(part1stacks[i][0]))
	}
	fmt.Fprintln(stdout, "")

	// Part 2 calculation
	part2stacks := cloneStacks(stacks)

	for i := range moves {
		m := moves[i]
		popped := make([]rune, m.count)
		remainder := make([]rune, len(part2stacks[m.from])-m.count)
		copy(popped, part2stacks[m.from][:m.count])
		copy(remainder, part2stacks[m.from][m.count:])
		part2stacks[m.from] = remainder
		part2stacks[m.to] = append(popped, part2stacks[m.to]...)
	}

	fmt.Fprint(stdout, "Top crates (part 2): ")
	for i := range stacks {
		fmt.Fprint(stdout, string(part2stacks[i][0]))
	}
	fmt.Fprintln(stdout, "")

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

func cloneStacks(original [][]rune) [][]rune {
	var result [][]rune
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(original)
	dec.Decode(&result)
	return result
}
