package main

import (
	"bytes"
	"os"
	"testing"
)

func TestRun_example(t *testing.T) {
	f, _ := os.Open("example-input.txt")
	var buf bytes.Buffer

	err := run([]string{}, f, &buf)
	if err != nil {
		t.Error("Received error:", err)
	}

	want := "Score (part 1): 15\nScore (part 2): 12\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}
