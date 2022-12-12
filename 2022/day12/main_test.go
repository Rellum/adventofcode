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

	want := "Answer (part 1): 31\nAnswer (part 2): Start at point (0,4) and take 29 steps\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}
