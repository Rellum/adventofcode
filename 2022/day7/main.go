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
	var wd string
	dirs := make(map[string]int64)

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return fmt.Errorf("unexpected empty line")
		}

		if line == "$ ls" {
			continue
		}

		if line == "$ cd /" {
			dirs["/"] += 0
			wd = "/"
			continue
		}

		if line == "$ cd .." {
			wd, _ = parent(wd)
			dirs[wd] += 0
			continue
		}

		if strings.HasPrefix(line, "$ cd ") {
			wd = wd + strings.TrimPrefix(line, "$ cd ") + "/"
			dirs[wd] += 0
			continue
		}

		if strings.HasPrefix(line, "dir ") {
			dirs[wd+strings.TrimPrefix(line, "dir ")+"/"] += 0
			continue
		}

		splits := strings.SplitN(line, " ", 2)

		i, err := strconv.ParseInt(splits[0], 10, 64)
		if err != nil {
			return fmt.Errorf("could not parse '%s' as number: %w", splits[0], err)
		}

		dir := wd
		var ok bool
		for {
			dirs[dir] += i
			dir, ok = parent(dir)
			if !ok {
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	var part1Sum int64
	for _, i := range dirs {
		if i > 100_000 {
			continue
		}

		part1Sum += i
	}

	fmt.Fprintln(stdout, "Part 1 sum:", part1Sum)

	return nil
}

func parent(path string) (string, bool) {
	if path == "/" {
		return "", false
	}
	splits := strings.Split(path, "/")
	return strings.Join(splits[:len(splits)-2], "/") + "/", true
}
