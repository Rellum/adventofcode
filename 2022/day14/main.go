package main

import (
	"bufio"
	"fmt"
	"image"
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
	viewport := image.Rect(500, 0, 500, 0)
	var paths [][]image.Point

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return fmt.Errorf("unexpected empty line")
		}

		points := strings.Split(line, " -> ")

		var path []image.Point
		for i := range points {
			splits := strings.Split(points[i], ",")
			if len(splits) != 2 {
				return fmt.Errorf("unexpected empty line")
			}

			x, err := strconv.Atoi(splits[0])
			if err != nil {
				return fmt.Errorf("parsing '%s' as int: %w", splits[0], err)
			}

			y, err := strconv.Atoi(splits[1])
			if err != nil {
				return fmt.Errorf("parsing '%s' as int: %w", splits[1], err)
			}

			path = append(path, image.Pt(x, y))

			if x < viewport.Min.X {
				viewport.Min.X = x
			}
			if viewport.Max.X <= x {
				viewport.Max.X = x + 1
			}
			if y < viewport.Min.Y {
				viewport.Min.Y = y
			}
			if viewport.Max.Y <= y {
				viewport.Max.Y = y + 1
			}
		}
		paths = append(paths, path)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	part1(stdout, viewport, paths)
	part2(stdout, viewport, paths)

	return nil
}

func part1(w io.Writer, viewport image.Rectangle, paths [][]image.Point) {
	s := createScan(viewport, paths)

	var count int
	for {
		resting := dropGrain(s)
		if !resting {
			break
		}
		count++
	}

	fmt.Fprintln(w, "Answer (part 1):", count)
}

func part2(w io.Writer, viewport image.Rectangle, paths [][]image.Point) {
	viewport.Min = viewport.Min.Add(image.Pt(-viewport.Dy(), 0))
	viewport.Max = viewport.Max.Add(image.Pt(viewport.Dy(), 2))
	paths = append(paths, []image.Point{image.Pt(viewport.Min.X, viewport.Max.Y-1), image.Pt(viewport.Max.X, viewport.Max.Y-1)})

	s := createScan(viewport, paths)

	count := 1
	for {
		resting := dropGrain(s)
		if !resting {
			break
		}
		count++
	}

	fmt.Fprintln(w, "Answer (part 2):", count)
}

type scan struct {
	viewport image.Rectangle
	source   image.Point
	tiles    []string
}

func createScan(viewport image.Rectangle, paths [][]image.Point) scan {
	res := scan{viewport: viewport, source: image.Pt(500, 0)}
	for i := viewport.Min.Y; i <= viewport.Max.Y; i++ {
		res.tiles = append(res.tiles, strings.Repeat(".", viewport.Max.X-viewport.Min.X+1))
	}
	set(res, res.source, '+')
	for i := 0; i < len(paths); i++ {
		path(res, paths[i])
	}
	return res
}

func path(s scan, path []image.Point) {
	for i := 0; i < len(path)-1; i++ {
		pathSegment(s, path[i], path[i+1])
	}
}

func pathSegment(s scan, a, b image.Point) {
	x1, x2 := a.X, b.X
	if x2 < x1 {
		x1, x2 = x2, x1
	}
	y1, y2 := a.Y, b.Y
	if y2 < y1 {
		y1, y2 = y2, y1
	}

	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			set(s, image.Pt(x, y), '#')
		}
	}
}

func dropGrain(s scan) bool {
	g := s.source
	for {
		g1 := dropGrainOneStep(s, g)
		if g1.Eq(s.source) {
			return false
		}
		if g1.Eq(g) {
			set(s, g, 'o')
			return true
		}
		if get(s, g1) == 0 {
			return false
		}
		g = g1
	}
}

func dropGrainOneStep(s scan, g image.Point) image.Point {
	offsets := []image.Point{
		image.Pt(0, 1),  // Down
		image.Pt(-1, 1), // Down and Left
		image.Pt(1, 1),  // Down and Right
	}
	for _, off := range offsets {
		g1 := g.Add(off)
		if !strings.ContainsRune("#o", get(s, g1)) {
			return g1
		}
	}

	return g
}

func get(s scan, pt image.Point) rune {
	if !pt.In(s.viewport) {
		return 0
	}

	x, y := pt.X-s.viewport.Min.X, pt.Y-s.viewport.Min.Y

	return rune(s.tiles[y][x])
}

func set(s scan, pt image.Point, c rune) {
	x, y := pt.X-s.viewport.Min.X, pt.Y-s.viewport.Min.Y
	s.tiles[y] = s.tiles[y][:x] + string(c) + s.tiles[y][x+1:]
}

func print(w io.Writer, s scan) {
	fmt.Fprintln(w, s.viewport)
	for i := range s.tiles {
		fmt.Fprintln(w, s.tiles[i])
	}
}
