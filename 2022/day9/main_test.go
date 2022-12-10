package main

import (
	"bytes"
	"os"
	"strconv"
	"testing"
)

func TestRun_example(t *testing.T) {
	f, _ := os.Open("example-input-1.txt")
	var buf bytes.Buffer

	err := run([]string{}, f, &buf)
	if err != nil {
		t.Error("Received error:", err)
	}

	want := "Answer (part 1): 13\nAnswer (part 2): 1\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}

func TestPulled(t *testing.T) {
	var head, tail coord

	tests := []struct {
		dir  rune
		tail coord
	}{
		// == R 4 ==
		{
			dir:  'R',
			tail: coord{},
		},
		{
			dir:  'R',
			tail: coord{1, 0},
		},
		{
			dir:  'R',
			tail: coord{2, 0},
		},
		{
			dir:  'R',
			tail: coord{3, 0},
		},
		// == U 4 ==
		{
			dir:  'U',
			tail: coord{3, 0},
		},
		{
			dir:  'U',
			tail: coord{4, 1},
		},
		{
			dir:  'U',
			tail: coord{4, 2},
		},
		{
			dir:  'U',
			tail: coord{4, 3},
		},
		// == L 3 ==
		{
			dir:  'L',
			tail: coord{4, 3},
		},
		{
			dir:  'L',
			tail: coord{3, 4},
		},
		{
			dir:  'L',
			tail: coord{2, 4},
		},
		// == D 1 ==
		{
			dir:  'D',
			tail: coord{2, 4},
		},
		// == R 4 ==
		{
			dir:  'R',
			tail: coord{2, 4},
		},
		{
			dir:  'R',
			tail: coord{2, 4},
		},
		{
			dir:  'R',
			tail: coord{3, 3},
		},
		{
			dir:  'R',
			tail: coord{4, 3},
		},
		// == D 1 ==
		{
			dir:  'D',
			tail: coord{4, 3},
		},
		// == L 5 ==
		{
			dir:  'L',
			tail: coord{4, 3},
		},
		{
			dir:  'L',
			tail: coord{4, 3},
		},
		{
			dir:  'L',
			tail: coord{3, 2},
		},
		{
			dir:  'L',
			tail: coord{2, 2},
		},
		{
			dir:  'L',
			tail: coord{1, 2},
		},
		// == R 2 ==
		{
			dir:  'R',
			tail: coord{1, 2},
		},
		{
			dir:  'R',
			tail: coord{1, 2},
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			head = move(head, test.dir)

			got := pulled(head, tail)

			if got != test.tail {
				t.Errorf("Incorrect output. Expected: %v\nGot:%v", test.tail, got)
			}
			tail = test.tail
		})
	}
}
