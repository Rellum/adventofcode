package main

import (
	"bytes"
	"crypto/md5"
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
	input, err := io.ReadAll(stdin)
	if err != nil {
		return fmt.Errorf("read all: %w", err)
	}

	input = bytes.TrimRight(input, "\n")

	answer5, answer6 := 0, 0

	i := 1
	for {
		digest := append(input, []byte(strconv.Itoa(i))...)
		hash := fmt.Sprintf("%x", md5.Sum(digest))

		if answer5 == 0 && strings.HasPrefix(hash, "00000") {
			answer5 = i
		}
		if strings.HasPrefix(hash, "000000") {
			answer6 = i
		}
		if answer5 > 0 && answer6 > 0 {
			break
		}

		i++
	}

	fmt.Fprintf(stdout, "Answer (5 zero prefix): %d\n", answer5)
	fmt.Fprintf(stdout, "Answer (6 zero prefix): %d\n", answer6)

	return nil
}
