package main

import (
	"bytes"
	"os"
	"strconv"
	"testing"
)

func TestRun_example(t *testing.T) {
	f, _ := os.Open("example-input.txt")
	var buf bytes.Buffer

	err := run([]string{}, f, &buf)
	if err != nil {
		t.Error("Received error:", err)
	}

	want := "Nice count: 2\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}

func TestIsNice(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "", want: false},
		{input: "ugknbfddgicrmopn", want: true},
		{input: "aaa", want: true},
		{input: "jchzalrnumimnmhp", want: false},
		{input: "haegwjzuvuyypxyu", want: false},
		{input: "dvszwmarrgswjxmb", want: false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := isNice(test.input)
			if got != test.want {
				t.Errorf("Incorrect output for '%s'.", test.input)
			}
		})
	}
}
