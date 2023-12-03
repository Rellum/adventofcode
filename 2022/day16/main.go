package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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

var rgx = regexp.MustCompile(`^Valve ([A-Z]+) has flow rate=([0-9]+); (tunnel leads to valve|tunnels lead to valves) ([A-Z, ]+)$`)

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	graph := make(map[string]valve)

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		matches := rgx.FindStringSubmatch(line)
		if matches == nil {
			return fmt.Errorf("line '%s' could not be parsed", line)
		}

		flowRate, err := strconv.Atoi(matches[2])
		if err != nil {
			return fmt.Errorf("could not parse '%s' as int: %w", matches[2], err)
		}

		graph[matches[1]] = valve{
			flowRate: flowRate,
			children: strings.Split(matches[4], ", "),
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	part1(stdout, graph)

	return nil
}

type valve struct {
	flowRate int
	children []string
}

type visit struct {
	minute   int
	location string
}

type state struct {
	name string
	prev *state
	visit
	on       []string
	released int
	flow     int
}

func part1(w io.Writer, graph map[string]valve) {

	fmt.Println(graph)
	res := search(graph)

	pretty(res, "result")

	fmt.Fprintln(w, "Answer (part 1):", res.released)
}

func search(graph map[string]valve) state {
	states := []state{{visit: visit{location: "AA"}}}
	visited := make(map[visit]state)
	var max state

	for len(states) > 0 {
		var s state
		s, states = pop(states)

		//if s.minute == 5 && s.released == 60 && len(s.on) == 2 {
		//	fmt.Println("debug minute 5", s, visited[s.visit])
		//}
		//
		//if s.minute == 6 && s.released == 93 {
		//	fmt.Println("debug minute 6", s)
		//}
		//
		//if s.minute == 5 && s.location == "II" && s.released == 180 {
		//	pretty(s)
		//}

		if prev, ok := visited[s.visit]; ok && s.flow <= prev.flow {
			continue
		}
		visited[s.visit] = s

		if s.minute == 30 {
			if max.released < s.released {
				max = s
			}
			continue
		}

		released := s.released + s.flow

		if graph[s.location].flowRate > 0 && !contains(s.on, s.location) {
			states = append(states, state{
				name: fmt.Sprintf("You open valve %s.", s.location),
				prev: &s,
				visit: visit{
					minute:   s.visit.minute + 1,
					location: s.visit.location,
				},
				on:       append(s.on, s.location),
				released: released,
				flow:     s.flow + graph[s.location].flowRate,
			})
		}

		for _, child := range graph[s.location].children {
			states = append(states, state{
				name: fmt.Sprintf("You move to valve %s.", child),
				prev: &s,
				visit: visit{
					minute:   s.visit.minute + 1,
					location: child,
				},
				on:       s.on,
				released: released,
				flow:     s.flow,
			})
		}
	}

	pretty(visited[visit{
		minute:   5,
		location: "BB",
	}], "debug minute 5")

	return max
}

func pop(a []state) (state, []state) {
	i := a[0]
	rem := append([]state{}, a[1:]...)
	return i, rem
}

//func search(graph map[string]valve, on []string, route []string, flow int) int {
//	if len(on)+len(route) == 30 {
//		return flow
//	}
//
//	current := route[len(route)-1]
//
//	var max int
//	if !contains(on, current) {
//		max = search(graph, append(on, current), route, flow+graph[current].flowRate) + flow
//	}
//
//	for _, child := range graph[current].children {
//		res := search(on, append(route, child), flow) + flow
//		if res > max {
//			max = res
//		}
//	}
//
//	fmt.Println(graph, on, route, flow, max)
//	//panic("foo")
//
//	return max
//}

func contains(sl []string, item string) bool {
	for i := range sl {
		if sl[i] == item {
			return true
		}
	}
	return false
}

func pretty(s state, prefix string) {
	for {
		fmt.Printf("== Minute %d == (%s)\n%v open, releasing %d (total %d)\n%s\n\n", s.minute, prefix, s.on, s.flow, s.released, s.name)
		if s.prev == nil {
			return
		}

		s = *s.prev
	}
}
