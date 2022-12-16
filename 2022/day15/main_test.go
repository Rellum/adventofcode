package main

import (
	"bytes"
	"os"
	"testing"
)

func TestRun_example(t *testing.T) {
	f, _ := os.Open("example-input.txt")
	var buf bytes.Buffer

	err := run([]string{"", "--row=10", "--square=20"}, f, &buf)
	if err != nil {
		t.Error("Received error:", err)
	}

	want := "Answer (part 1): 26\nAnswer (part 2): 56000011\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}
