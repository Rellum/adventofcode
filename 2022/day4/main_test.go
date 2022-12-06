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

	want := "Count fully contained (part 1): 2\nCount overlapping (part 2): 4\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}

func TestAssignment_overlaps(t *testing.T) {
	tests := []struct {
		a, b assignment
		want bool
	}{
		{
			a: assignment{1, 3},
			b: assignment{4, 6},
		},
		{
			a: assignment{4, 6},
			b: assignment{1, 3},
		},
		{
			a:    assignment{1, 4},
			b:    assignment{4, 6},
			want: true,
		},
		{
			a:    assignment{4, 6},
			b:    assignment{1, 4},
			want: true,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test), func(t *testing.T) {
			if test.a.overlaps(&test.b) != test.want {
				t.Fail()
			}
		})
	}
}
