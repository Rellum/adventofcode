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
		{input: "", want: "Year one house count: 1\nYear two house count: 1"},
		{input: ">", want: "Year one house count: 2\nYear two house count: 2"},
		{input: "^v", want: "Year one house count: 2\nYear two house count: 3"},
		{input: "^>v<", want: "Year one house count: 4\nYear two house count: 3"},
		{input: "^v^v^v^v^v", want: "Year one house count: 2\nYear two house count: 11"},
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
