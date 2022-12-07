package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	var nice1Count, nice2Count int

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return fmt.Errorf("unexpected empty line")
		}

		if isNice1(line) {
			nice1Count++
		}

		if isNice2(line) {
			nice2Count++
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Fprintf(stdout, "Nice count (part 1): %d\n", nice1Count)
	fmt.Fprintf(stdout, "Nice count (part 2): %d\n", nice2Count)

	return nil
}

func isNice1(s string) bool {
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

func isNice2(s string) bool {
	pairs := map[string]struct{}{}
	var doublePair, repeatedLetter bool

	splits := strings.Split(s, "")
	for i := range splits {
		if i > 2 {
			pairs[splits[i-3]+splits[i-2]] = struct{}{}
		}

		if i < 3 {
			continue
		}

		if splits[i-2] == splits[i] {
			repeatedLetter = true
		}

		if _, ok := pairs[splits[i-1]+splits[i]]; ok {
			doublePair = true
		}
	}

	return repeatedLetter && doublePair
}
