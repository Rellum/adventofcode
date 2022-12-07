package main

import (
	"bytes"
	"os"
	"testing"
)

func TestRun_example(t *testing.T) {
	f, _ := os.Open("../example-input.txt")
	var buf bytes.Buffer

	err := run([]string{}, f, &buf)
	if err != nil {
		t.Error("Received error:", err)
	}

	want, err := os.ReadFile("expected.txt")
	if err != nil {
		t.Errorf("read file: %s", err)
	}

	got := buf.String()
	if got != string(want) {
		t.Errorf("Incorrect output:\n>>>%s<<<\n\nWant:\n>>>%s<<<", got, want)
	}
}
