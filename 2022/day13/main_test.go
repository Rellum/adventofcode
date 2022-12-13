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

	want := "Answer (part 1): 13\nAnswer (part 2): 140\n"
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
		{
			input: "[[[6,1,[],[7]]],[[3],1,10,1,6]]\n" +
				"[[[[6,10,7],1,[8],[3,9,2,4,0],8],2,[[8,1,6,9],[9,10,9,7,0],[1],[4,10,6]]],[[4,6,[0,2],3,10],[]]]\n",
			want: "Answer (part 1): 1\nAnswer (part 2): 2",
		},
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
