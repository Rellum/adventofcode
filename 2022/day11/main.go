package main

import (
	"bufio"
	"flag"
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
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		relief = flags.Bool("relief", false, "whether or not the worry is diminished by relief")
		rounds = flags.Int("rounds", 1, "number of rounds to play")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	var monkeys []monkey

	scanner := bufio.NewScanner(stdin)
	var m monkey
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "Monkey ") {
			m = monkey{}
			continue
		}

		line = strings.TrimLeft(line, " ")

		if s, ok := matchPrefix(line, "Starting items: "); ok {
			splits := strings.Split(s, ", ")
			for i := range splits {
				item, err := strconv.Atoi(splits[i])
				if err != nil {
					return fmt.Errorf("could not parse '%s' as int: %w", splits[i], err)
				}
				m.items = append(m.items, item)
			}
			continue
		}

		if s, ok := matchPrefix(line, "Operation: new = "); ok {
			if s == "old * old" {
				m.operation = square
				continue
			}

			if s, ok := matchPrefix(s, "old + "); ok {
				addend, err := strconv.Atoi(s)
				if err != nil {
					return fmt.Errorf("could not parse '%s' as int: %w", s, err)
				}

				m.operation = addFunc(addend)
				continue
			}

			if s, ok := matchPrefix(s, "old * "); ok {
				factor, err := strconv.Atoi(s)
				if err != nil {
					return fmt.Errorf("could not parse '%s' as int: %w", s, err)
				}

				m.operation = multiplyFunc(factor)
				continue
			}

			return fmt.Errorf("could not parse operation '%s'", s)
		}

		if s, ok := matchPrefix(line, "Test: divisible by "); ok {
			testMod, err := strconv.Atoi(s)
			if err != nil {
				return fmt.Errorf("could not parse '%s' as int: %w", s, err)
			}
			m.testMod = testMod
			continue
		}

		if s, ok := matchPrefix(line, "If true: throw to monkey "); ok {
			recipient, err := strconv.Atoi(s)
			if err != nil {
				return fmt.Errorf("could not parse '%s' as int: %w", s, err)
			}
			m.passTrue = recipient
			continue
		}

		if s, ok := matchPrefix(line, "If false: throw to monkey "); ok {
			recipient, err := strconv.Atoi(s)
			if err != nil {
				return fmt.Errorf("could not parse '%s' as int: %w", s, err)
			}
			m.passFalse = recipient

			monkeys = append(monkeys, m)

			continue
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	mostActive := keepAway(monkeys, *rounds, *relief)
	fmt.Fprintf(stdout, "Answer: %d\n", mostActive[0]*mostActive[1])

	return nil
}

func matchPrefix(s, prefix string) (string, bool) {
	if !strings.HasPrefix(s, prefix) {
		return "", false
	}
	return strings.TrimPrefix(s, prefix), true
}

type monkey struct {
	items               []int
	operation           func(int) int
	inspections         int
	testMod             int
	passTrue, passFalse int
}

func square(old int) int {
	return old * old
}

func addFunc(addend int) func(int) int {
	return func(old int) int {
		return old + addend
	}
}

func multiplyFunc(factor int) func(int) int {
	return func(old int) int {
		return old * factor
	}
}

func keepAway(monkeys []monkey, rounds int, relief bool) [2]int {
	for i := 0; i < rounds; i++ {
		round(monkeys, relief)
	}

	var first, second int
	for i := 0; i < len(monkeys); i++ {
		if monkeys[i].inspections > first {
			second = first
			first = monkeys[i].inspections
			continue
		}
		if monkeys[i].inspections > second {
			second = monkeys[i].inspections
		}
	}

	return [...]int{first, second}
}

func round(monkeys []monkey, relief bool) {
	lcm := 1
	for i := 0; i < len(monkeys); i++ {
		lcm *= monkeys[i].testMod
	}

	for i := 0; i < len(monkeys); i++ {
		m := &monkeys[i]
		var item int
		var itemCount = len(m.items)
		for j := 0; j < itemCount; j++ {
			m.inspections++

			item, m.items = pop(m.items)

			item = m.operation(item)

			if relief {
				item /= 3
			}

			item %= lcm

			recipient := m.passFalse
			if item%m.testMod == 0 {
				recipient = m.passTrue
			}

			monkeys[recipient].items = append(monkeys[recipient].items, item)
		}
	}
}

func pop(a []int) (int, []int) {
	i := a[0]
	rem := append([]int{}, a[1:]...)
	return i, rem
}
