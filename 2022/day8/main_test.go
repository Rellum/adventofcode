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

	want := `vvvvv
vvv.v
vv.vv
v.v.v
vvvvv
Answer (part 1): 21
Answer (part 2): 8
`
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}

func TestScenicScore(t *testing.T) {
	grid := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	tests := []struct {
		row, col, score int
	}{
		{
			row:   1,
			col:   2,
			score: 4,
		},
		{
			row:   3,
			col:   2,
			score: 8,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("r%dc%d", test.row, test.col), func(t *testing.T) {
			got := scenicScore(grid, test.row, test.col)
			if got != test.score {
				t.Errorf("Incorrect output. Got %d; Expected %d", got, test.score)
			}
		})
	}
}
