package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
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

var rgx = regexp.MustCompile(`^Sensor at x=([-0-9]+), y=([-0-9]+): closest beacon is at x=([-0-9]+), y=([-0-9]+)$`)

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		row    = flags.Int("row", 10, "which row to count")
		square = flags.Int("square", 20, "size of square to search")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	var viewport *image.Rectangle
	var sensors []sensor

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		matches := rgx.FindStringSubmatch(line)
		if matches == nil {
			return fmt.Errorf("line could not be parsed: '%s'", line)
		}

		var ints [4]int
		var err error
		for i := 1; i <= 4; i++ {
			ints[i-1], err = strconv.Atoi(matches[i])
			if err != nil {
				return fmt.Errorf("parsing '%s' as int: %w", matches[i], err)
			}
		}
		s := makeSensor(image.Pt(ints[0], ints[1]), image.Pt(ints[2], ints[3]))
		sensors = append(sensors, s)

		offsets := []image.Point{
			image.Pt(0, -1), // Up
			image.Pt(1, 0),  // Right
			image.Pt(0, 1),  // Down
			image.Pt(-1, 0), // Left
		}
		for _, off := range offsets {
			viewport = zoom(viewport, s.pos.Add(off.Mul(s.dist)))
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	part1(stdout, *viewport, sensors, *row)
	part2(stdout, image.Rect(0, 0, *square, *square), sensors)

	return nil
}

func part1(w io.Writer, viewport image.Rectangle, sensors []sensor, row int) {
	var count int
	for x := viewport.Min.X; x < viewport.Max.X; x++ {
		p := image.Pt(x, row)
		var eliminated bool
		for i := 0; i < len(sensors); i++ {
			s := sensors[i]
			if p == s.pos || p == s.closest {
				i = len(sensors)
				break
			}
			d := dist(p, s.pos)
			if d <= s.dist {
				eliminated = true
				i = len(sensors)
				break
			}
		}
		if eliminated {
			count++
		}
	}

	fmt.Fprintln(w, "Answer (part 1):", count)
}

func part2(w io.Writer, viewport image.Rectangle, sensors []sensor) {
	var p image.Point
	for y := viewport.Min.Y; y < viewport.Max.Y; y++ {
		for x := viewport.Min.X; x < viewport.Max.X; x++ {
			p = image.Pt(x, y)
			var eliminated bool
			for i := 0; i < len(sensors); i++ {
				s := sensors[i]
				if p == s.pos || p == s.closest {
					eliminated = true
					break
				}
				d := dist(p, s.pos)
				if d <= s.dist {
					eliminated = true
					x = s.pos.X + s.dist - abs(s.pos.Y-y)

					break
				}
			}
			if !eliminated {
				fmt.Fprintln(w, "Answer (part 2):", p.X*4_000_000+p.Y)
				return
			}
		}
	}
}

type sensor struct {
	pos, closest image.Point
	dist         int
}

func makeSensor(pos, closest image.Point) sensor {
	return sensor{
		pos:     pos,
		closest: closest,
		dist:    dist(pos, closest),
	}
}

func dist(a, b image.Point) int {
	var res int
	if a.X < b.X {
		res += b.X - a.X
	} else {
		res += a.X - b.X
	}
	if a.Y < b.Y {
		res += b.Y - a.Y
	} else {
		res += a.Y - b.Y
	}
	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func zoom(r *image.Rectangle, pt image.Point) *image.Rectangle {
	if r == nil {
		return &image.Rectangle{Min: pt, Max: pt}
	}

	if pt.X < r.Min.X {
		r.Min.X = pt.X
	}
	if r.Max.X <= pt.X {
		r.Max.X = pt.X + 1
	}
	if pt.Y < r.Min.Y {
		r.Min.Y = pt.Y
	}
	if r.Max.Y <= pt.Y {
		r.Max.Y = pt.Y + 1
	}

	return r
}
