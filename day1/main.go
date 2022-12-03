package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	calories := []int64{0}
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			calories = append(calories, 0)
			continue
		}

		i, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse '%s': %w", line, err)
		}

		calories[len(calories)-1] += i
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	sort.SliceStable(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	fmt.Fprintln(stdout, "Max:", calories[0])
	fmt.Fprintln(stdout, "Sum of top 3:", calories[0]+calories[1]+calories[2])

	return nil
}
