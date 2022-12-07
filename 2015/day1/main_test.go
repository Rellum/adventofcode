package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "", want: "Floor: 0\n"},
		{input: "(())", want: "Floor: 0\n"},
		{input: "()()", want: "Floor: 0\n"},
		{input: "(((", want: "Floor: 3\n"},
		{input: "(()(()(", want: "Floor: 3\n"},
		{input: "))(((((", want: "Floor: 3\nEnters basement at: 1\n"},
		{input: "())", want: "Floor: -1\nEnters basement at: 3\n"},
		{input: "))(", want: "Floor: -1\nEnters basement at: 1\n"},
		{input: ")))", want: "Floor: -3\nEnters basement at: 1\n"},
		{input: ")())())", want: "Floor: -3\nEnters basement at: 1\n"},
		{input: ")", want: "Floor: -1\nEnters basement at: 1\n"},
		{input: "()())", want: "Floor: -1\nEnters basement at: 5\n"},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var buf bytes.Buffer

			err := run([]string{}, strings.NewReader(test.input), &buf)
			if err != nil {
				t.Error("run:", err)
			}

			got := buf.String()
			if got != test.want {
				t.Errorf("Incorrect output for '%s'. Expected:\n>>>%s<<<\nGot:\n>>>%s<<<", test.input, test.want, got)
			}
		})
	}
}
