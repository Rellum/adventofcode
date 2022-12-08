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
	bright := map[coord]int{}

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			grid[coord{x, y}] = false
			bright[coord{x, y}] = 0
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
				c := coord{x, y}
				switch matches[1] {
				case "turn on":
					grid[c] = true
					bright[c]++
				case "turn off":
					grid[c] = false
					if bright[c] > 0 {
						bright[c]--
					}
				case "toggle":
					grid[c] = !grid[c]
					bright[c] += 2
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	var count int
	var totalBrightness int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			c := coord{x, y}
			if grid[c] {
				count++
			}
			totalBrightness += bright[c]
		}
	}

	fmt.Fprintf(stdout, "Answer (part1): %d\n", count)
	fmt.Fprintf(stdout, "Answer (part2): %d\n", totalBrightness)

	return nil
}

type coord struct{ x, y int }
