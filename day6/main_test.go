package main

import (
	"bytes"
	"fmt"
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

	want := "First marker after character 7\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  7,
		},
		{
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  5,
		},
		{
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  6,
		},
		{
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  10,
		},
		{
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  11,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := part1(test.input)
			if err != nil {
				t.Errorf("error: %v", err)
			}

			if got != test.want {
				t.Errorf("want: %d; got: %d", test.want, got)
			}
		})
	}
}
