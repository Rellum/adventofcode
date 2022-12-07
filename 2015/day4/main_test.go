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
		{input: "abcdef", want: "Answer (5 zero prefix): 609043\nAnswer (6 zero prefix): 6742839"},
		{input: "pqrstuv", want: "Answer (5 zero prefix): 1048970\nAnswer (6 zero prefix): 5714438"},
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
