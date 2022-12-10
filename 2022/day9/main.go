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
	var head, tail coord
	visited1 := map[coord]struct{}{}
	snake := [9]coord{}
	visited2 := map[coord]struct{}{}

	// Scan lines
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return fmt.Errorf("unexpected empty line")
		}

		splits := strings.Split(line, " ")

		if len(splits) != 2 {
			return fmt.Errorf("'%s' could not be parsed into direction and number", line)
		}

		count, err := strconv.Atoi(splits[1])
		if err != nil {
			return fmt.Errorf("'%s' could not be parsed as a number: %w", splits[1], err)
		}

		runes := []rune(splits[0])
		if len(runes) != 1 {
			return fmt.Errorf("'%s' is not a single character", splits[0])
		}
		dir := runes[0]

		for i := 0; i < count; i++ {
			head = move(head, dir)
			tail = pulled(head, tail)

			prevAfter := head
			for j := 0; j < len(snake); j++ {
				cur := snake[j]
				snake[j] = pulled(prevAfter, cur)

				prevAfter = snake[j]
			}

			visited1[tail] = struct{}{}
			visited2[snake[len(snake)-1]] = struct{}{}
		}

		//fmt.Println(line)
		//print([]coord{head, tail}, visited1)
		//print(append([]coord{head}, snake[:]...), visited2)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Fprintf(stdout, "Answer (part 1): %d\n", len(visited1))
	fmt.Fprintf(stdout, "Answer (part 2): %d\n", len(visited2))

	return nil
}

func pulled(headPos, tailPos coord) coord {
	var xDist, yDist int
	xDist = (headPos.x - tailPos.x) * (headPos.x - tailPos.x)
	yDist = (headPos.y - tailPos.y) * (headPos.y - tailPos.y)

	res := tailPos

	if yDist < 2 && xDist <= 2 {
		return res
	}

	if yDist > 1 && xDist == 0 {
		if headPos.y > tailPos.y {
			res.y = tailPos.y + 1
		} else {
			res.y = tailPos.y - 1
		}
	} else if xDist > 1 && yDist == 0 {
		if headPos.x > tailPos.x {
			res.x = tailPos.x + 1
		} else {
			res.x = tailPos.x - 1
		}
	} else {
		if headPos.y > tailPos.y {
			res.y = tailPos.y + 1
		} else {
			res.y = tailPos.y - 1
		}
		if headPos.x > tailPos.x {
			res.x = tailPos.x + 1
		} else {
			res.x = tailPos.x - 1
		}
	}

	return res
}

type coord struct {
	x, y int
}

func move(c coord, d rune) coord {
	switch d {
	case 'L':
		return coord{c.x - 1, c.y}
	case 'R':
		return coord{c.x + 1, c.y}
	case 'U':
		return coord{c.x, c.y + 1}
	case 'D':
		return coord{c.x, c.y - 1}
	default:
		panic(fmt.Sprintf("'%s' is not a valid direction", string(d)))
	}
}

func print(s []coord, visited map[coord]struct{}) {
	m := make(map[coord]int)
	for i := len(s) - 1; i >= 0; i-- {
		m[s[i]] = i
	}

	for y := 17; y > -6; y-- {
		for x := -15; x < 20; x++ {
			c := coord{x, y}
			if i, ok := m[c]; ok {
				fmt.Print(i)
				continue
			}
			if _, ok := visited[c]; ok {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
	fmt.Println("===")
}
