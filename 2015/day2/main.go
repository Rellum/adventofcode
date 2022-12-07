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
	var totalPaper int
	var totalRibbon int

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return fmt.Errorf("unexpected empty line")
		}

		paper, ribbon, err := wrapOne(line)
		if err != nil {
			return fmt.Errorf("wrapping '%s': %w", line, err)
		}

		totalPaper += paper
		totalRibbon += ribbon
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Fprintf(stdout, "%d square feet of wrapping paper\n", totalPaper)
	fmt.Fprintf(stdout, "%d feet of ribbon\n", totalRibbon)

	return nil
}

func wrapOne(dims string) (int, int, error) {
	splits := strings.SplitN(dims, "x", 3)

	var sides [3]int
	var err error
	for i := 0; i < 3; i++ {
		sides[i], err = strconv.Atoi(splits[i])
		if err != nil {
			return 0, 0, fmt.Errorf("convert '%s' to a number: %w\n", splits[i], err)
		}
	}

	var total, smallestArea, smallestPerimeter int
	bow := sides[0] * sides[1] * sides[2]
	for _, combo := range [3][2]int{{0, 1}, {0, 2}, {1, 2}} {
		area := sides[combo[0]] * sides[combo[1]]
		perimeter := sides[combo[0]]*2 + sides[combo[1]]*2

		total += 2 * area

		if smallestArea == 0 || area < smallestArea {
			smallestArea = area
		}

		if smallestPerimeter == 0 || perimeter < smallestPerimeter {
			smallestPerimeter = perimeter
		}
	}

	return total + smallestArea, bow + smallestPerimeter, nil
}
