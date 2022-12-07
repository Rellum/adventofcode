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
	var floor int

	r := bufio.NewReader(stdin)
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return fmt.Errorf("read rune: %w", err)
			}
		}

		switch c {
		case '(':
			floor++
		case ')':
			floor--
		case '\n':
			break
		default:
			return fmt.Errorf("unexpected character: '%s'", string(c))
		}
	}

	fmt.Fprintf(stdout, "Floor: %d\n", floor)

	return nil
}
