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

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	locY1 := coord{}
	locY2S1 := coord{}
	locY2S2 := coord{}
	visitedYear1 := map[coord]int{locY1: 1}
	visitedYear2 := map[coord]int{coord{}: 2}

	r := bufio.NewReader(stdin)
	var i int
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return fmt.Errorf("read rune: %w", err)
			}
		}

		locY1, err = move(locY1, c)
		if err != nil && !errors.Is(err, io.EOF) {
			return fmt.Errorf("read rune: %w", err)
		}

		if i%2 == 0 {
			locY2S1, err = move(locY2S1, c)
			if err != nil && !errors.Is(err, io.EOF) {
				return fmt.Errorf("read rune: %w", err)
			}
		} else {
			locY2S2, err = move(locY2S2, c)
			if err != nil && !errors.Is(err, io.EOF) {
				return fmt.Errorf("read rune: %w", err)
			}
		}

		visitedYear1[locY1]++
		visitedYear2[locY2S1]++
		visitedYear2[locY2S2]++
		i++
	}

	fmt.Fprintf(stdout, "Year one house count: %d\n", len(visitedYear1))
	fmt.Fprintf(stdout, "Year two house count: %d\n", len(visitedYear2))

	return nil
}

type coord struct {
	x, y int
}

func move(from coord, dir rune) (coord, error) {
	switch dir {
	case '<':
		return coord{from.x - 1, from.y}, nil
	case '>':
		return coord{from.x + 1, from.y}, nil
	case '^':
		return coord{from.x, from.y + 1}, nil
	case 'v':
		return coord{from.x, from.y - 1}, nil
	case '\n':
		return coord{}, io.EOF
	default:
		return coord{}, fmt.Errorf("unexpected character: '%s'", string(dir))
	}
}
