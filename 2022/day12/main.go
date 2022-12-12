package main

import (
	"bufio"
	"fmt"
	"github.com/fzipp/astar"
	"image"
	"io"
	"math"
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
	var start, end image.Point
	var valleys []image.Point
	var hm heightmap

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		hm = append(hm, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	for y := 0; y < len(hm); y++ {
		for x := 0; x < len(hm[y]); x++ {
			if hm[y][x] == 'S' {
				start = image.Pt(x, y)
				continue
			}
			if hm[y][x] == 'E' {
				end = image.Pt(x, y)
				continue
			}
			if hm[y][x] == 'a' {
				valleys = append(valleys, image.Pt(x, y))
			}
		}
	}

	fmt.Fprintf(stdout, "Answer (part 1): %d\n", len(search(hm, start, end))-1)

	var shortest astar.Path[image.Point]
	for i := range valleys {
		path := search(hm, valleys[i], end)
		if len(path) == 0 {
			continue
		}
		if len(shortest) != 0 && len(shortest) <= len(path) {
			continue
		}

		shortest = path
	}

	fmt.Fprintf(stdout, "Answer (part 2): Start at point %v and take %d steps\n", shortest[0], len(shortest)-1)

	return nil
}

func search(hm heightmap, start, end image.Point) astar.Path[image.Point] {
	return astar.FindPath[image.Point](hm, start, end, distance, distance)
}

func distance(p, q image.Point) float64 {
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

type heightmap []string

func (h heightmap) Neighbours(p image.Point) []image.Point {
	offsets := []image.Point{
		image.Pt(0, -1), // Up
		image.Pt(1, 0),  // Right
		image.Pt(0, 1),  // Down
		image.Pt(-1, 0), // Left
	}
	res := make([]image.Point, 0, 4)
	for _, off := range offsets {
		q := p.Add(off)
		if possible(h, p, q) {
			res = append(res, q)
		}
	}
	return res
}

func possible(h heightmap, from, to image.Point) bool {
	if !withinBounds(h, to) {
		return false
	}

	fromHeight, toHeight := h[from.Y][from.X], h[to.Y][to.X]

	if fromHeight == 'S' {
		fromHeight = 'a'
	}

	if toHeight == 'E' {
		toHeight = 'z'
	}

	maxHeight := fromHeight + 1
	if maxHeight > 'z' {
		maxHeight = 'z'
	}

	return toHeight <= maxHeight
}

func withinBounds(h heightmap, p image.Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(h) && p.X < len(h[p.Y])
}
