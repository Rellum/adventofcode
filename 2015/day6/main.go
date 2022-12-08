package main

import (
	"bufio"
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

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

var rgx = regexp.MustCompile(`^(turn on|turn off|toggle) (\d{1,3}),(\d{1,3}) through (\d{1,3}),(\d{1,3})$`)

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	grid := map[coord]bool{}

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			grid[coord{x, y}] = false
		}
	}

	// Scan lines
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		matches := rgx.FindStringSubmatch(line)
		if matches == nil {
			return fmt.Errorf("instruction '%s' could not be parsed")
		}

		x1, _ := strconv.Atoi(matches[2])
		y1, _ := strconv.Atoi(matches[3])
		x2, _ := strconv.Atoi(matches[4])
		y2, _ := strconv.Atoi(matches[5])

		if x1 > x2 {
			old := x1
			x1 = x2
			x2 = old
		}

		if y1 > y2 {
			old := y1
			y1 = y2
			y2 = old
		}

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				switch matches[1] {
				case "turn on":
					grid[coord{x, y}] = true
				case "turn off":
					grid[coord{x, y}] = false
				case "toggle":
					grid[coord{x, y}] = !grid[coord{x, y}]
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	var count int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[coord{x, y}] {
				count++
			}
		}
	}

	fmt.Fprintf(stdout, "Answer: %d\n", count)

	return nil
}

type coord struct{ x, y int }
