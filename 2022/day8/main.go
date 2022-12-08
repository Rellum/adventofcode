package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	var heights [][]int
	var visible [][]bool

	// Scan lines
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return fmt.Errorf("unexpected empty line")
		}

		heights = append(heights, []int{})
		visible = append(visible, []bool{})
		row := len(heights) - 1

		for _, r := range line {
			i, err := strconv.Atoi(string(r))
			if err != nil {
				return fmt.Errorf("integer conversion of '%s': %w", string(r), err)
			}

			heights[row] = append(heights[row], i)
			visible[row] = append(visible[row], false)
		}

		row++
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	var highest int

	// look from left
	for row := 0; row < len(heights); row++ {
		highest = -1
		for col := 0; col < len(heights[0]); col++ {
			if height := heights[row][col]; height > highest {
				visible[row][col] = true
				highest = height
			}
		}
	}

	// look from right
	for row := 0; row < len(heights); row++ {
		highest = -1
		for col := len(heights[0]) - 1; col >= 0; col-- {
			if height := heights[row][col]; height > highest {
				visible[row][col] = true
				highest = height
			}
		}
	}

	// look from top
	for col := 0; col < len(heights[0]); col++ {
		highest = -1
		for row := 0; row < len(heights); row++ {
			if height := heights[row][col]; height > highest {
				visible[row][col] = true
				highest = height
			}
		}
	}

	// look from bottom
	for col := 0; col < len(heights[0]); col++ {
		highest = -1
		for row := len(heights[0]) - 1; row >= 0; row-- {
			if height := heights[row][col]; height > highest {
				visible[row][col] = true
				highest = height
			}
		}
	}

	var count int
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[0]); col++ {
			if visible[row][col] {
				fmt.Fprintf(stdout, "v")
				count++
			} else {
				fmt.Fprintf(stdout, ".")
			}
		}
		fmt.Fprintf(stdout, "\n")
	}

	fmt.Fprintf(stdout, "Answer: %v\n", count)

	return nil
}
