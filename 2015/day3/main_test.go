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
		{input: "", want: "1"},
		{input: ">", want: "2"},
		{input: "^>v<", want: "4"},
		{input: "^v^v^v^v^v", want: "2"},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var buf bytes.Buffer

			err := run([]string{}, strings.NewReader(test.input), &buf)
			if err != nil {
				t.Error("run:", err)
			}

			got := strings.TrimRight(buf.String(), "\n")
			if got != test.want {
				t.Errorf("Incorrect output for '%s'. Expected:\n>>>%s<<<\nGot:\n>>>%s<<<", test.input, test.want, got)
			}
		})
	}
}
