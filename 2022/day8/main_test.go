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

	want := `vvvvv
vvv.v
vv.vv
v.v.v
vvvvv
Answer: 21
`
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}
