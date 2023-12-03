package main

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestRun_example(t *testing.T) {
	f, _ := os.Open("example-input.txt")
	var buf bytes.Buffer

	err := run([]string{}, f, &buf)
	if err != nil {
		t.Error("Received error:", err)
	}

	want := "day {{cookiecutter.day}} answer\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}

func Test(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "", want: "day {{cookiecutter.day}} answer"},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var buf bytes.Buffer

			err := run([]string{}, strings.NewReader(test.input), &buf)
			if err != nil {
				t.Error("run:", err)
			}

			want := test.want + "\n"
			got := buf.String()
			if got != want {
				t.Errorf("Incorrect output for '%s'. Expected:\n>>>%s<<<\nGot:\n>>>%s<<<", test.input, want, got)
			}
		})
	}
}