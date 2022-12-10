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

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	hist := []int{1}

	// Scan lines
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return fmt.Errorf("unexpected empty line")
		}

		if line == "noop" {
			hist = append(hist, hist[len(hist)-1])
		}

		if strings.HasPrefix(line, "addx ") {
			s := strings.TrimPrefix(line, "addx ")
			addend, err := strconv.Atoi(s)
			if err != nil {
				return fmt.Errorf("parsing '%s' as number: %w", s, err)
			}

			hist = append(hist, hist[len(hist)-1])
			hist = append(hist, hist[len(hist)-1]+addend)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	if len(hist) < 220 {
		fmt.Fprintf(stdout, "History: %v\n", hist)
	}

	var sum int
	for i := 20; i <= 220; i += 40 {
		sum += i * hist[i-1]
	}

	fmt.Fprintf(stdout, "Answer: %#v\n", sum)

	for i := 0; i < 40*6; i++ {
		xPos := (i) % 40

		if xPos == 0 {
			fmt.Fprint(stdout, "\n")
		}

		dist := xPos - hist[i]
		if -1 <= dist && dist <= 1 {
			fmt.Fprint(stdout, "#")
		} else {
			fmt.Fprint(stdout, ".")
		}
	}

	fmt.Fprint(stdout, "\n")

	return nil
}
