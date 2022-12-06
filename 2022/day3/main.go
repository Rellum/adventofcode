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

	groupSize = 3
)

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	var part1sum int
	var part2sum int
	var i int // elf number (starts at 0)
	var j int // elf number in group (starts at 0)
	var bitfields [52]uint8

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return fmt.Errorf("unexpected empty line")
		}

		if len(line)%2 != 0 {
			return fmt.Errorf("odd number of items not allowed")
		}

		j = (i) % groupSize

		for _, c := range line {
			prio, err := priority(c)
			if err != nil {
				return fmt.Errorf("error in line '%s': %w", line, err)
			}

			bitfields[prio-1] = bitfields[prio-1] | 1<<j
		}

		if j == 2 {
			for j, v := range bitfields {
				if v == 7 {
					part2sum += int(j) + 1
					break
				}
			}
			bitfields = [52]uint8{}
		}
		i++

		compartment1, compartment2 := line[:len(line)/2], line[len(line)/2:]

		m := make(map[rune]struct{})
		for _, c := range compartment1 {
			m[c] = struct{}{}
		}
		for _, c := range compartment2 {
			if _, ok := m[c]; ok {
				prio, err := priority(c)
				if err != nil {
					return fmt.Errorf("error in line '%s': %w", line, err)
				}

				part1sum += prio
				break
			}
		}

	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Fprintln(stdout, "Sum (part 1):", part1sum)
	fmt.Fprintln(stdout, "Sum (part 2):", part2sum)

	return nil
}

func priority(c rune) (int, error) {
	if 'a' <= c && c <= 'z' {
		return int(c-'a') + 1, nil
	}

	if 'A' <= c && c <= 'Z' {
		return int(c-'A') + 27, nil
	}

	return 0, errors.New("character out of range")
}
