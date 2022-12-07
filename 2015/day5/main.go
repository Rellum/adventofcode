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
	var niceCount int

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return fmt.Errorf("unexpected empty line")
		}

		if isNice(line) {
			niceCount++
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Fprintf(stdout, "Nice count: %d\n", niceCount)

	return nil
}

func isNice(s string) bool {
	var vowels int
	var double bool

	var last rune
	for _, c := range s {
		for _, pair := range []string{"ab", "cd", "pq", "xy"} {
			if last == int32(pair[0]) && c == int32(pair[1]) {
				return false
			}
		}

		if c == last {
			double = true
		}

		for _, vowel := range "aeiou" {
			if c == vowel {
				vowels++
				break
			}
		}

		last = c
	}

	return double && vowels >= 3
}
