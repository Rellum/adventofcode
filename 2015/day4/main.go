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

	answer := 1
	for {
		digest := append(input, []byte(strconv.Itoa(answer))...)
		hash := fmt.Sprintf("%x", md5.Sum(digest))
		if strings.HasPrefix(hash, "00000") {
			break
		}
		answer++
	}

	fmt.Fprintf(stdout, "Answer: %d\n", answer)

	return nil
}
