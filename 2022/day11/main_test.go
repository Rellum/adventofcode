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

	err := run([]string{"", "--relief=true", "--rounds=20"}, f, &buf)
	if err != nil {
		t.Error("Received error:", err)
	}

	want := "Answer: 10605\n"
	got := buf.String()
	if got != want {
		t.Error("Incorrect output:", got)
	}
}

func TestKeepAway(t *testing.T) {
	tests := []struct {
		rounds int
		want   [2]int
		relief bool
	}{
		{relief: true, rounds: 20, want: [2]int{105, 101}},
		{rounds: 1, want: [2]int{6, 4}},
		{rounds: 20, want: [2]int{103, 99}},
		{rounds: 1_000, want: [2]int{5_204, 5_192}},
		{rounds: 10_000, want: [2]int{52_166, 52_013}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("relief_%t_rounds_%d", test.relief, test.rounds), func(t *testing.T) {
			monkeys := exampleMonkeys()
			got := keepAway(monkeys, test.rounds, test.relief)

			if got[0] != test.want[0] || got[1] != test.want[1] {
				t.Errorf("Expected:\n>>>%v<<<\nGot:\n>>>%v<<<", test.want, got)
			}
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		rounds int
		want   []int
		relief bool
	}{
		{relief: true, rounds: 20, want: []int{101, 95, 7, 105}},
		{rounds: 1, want: []int{2, 4, 3, 6}},
		{rounds: 20, want: []int{99, 97, 8, 103}},
		{rounds: 1_000, want: []int{5_204, 4_792, 199, 5_192}},
		{rounds: 10_000, want: []int{52_166, 47_830, 1_938, 52_013}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("relief_%t_rounds_%d", test.relief, test.rounds), func(t *testing.T) {
			monkeys := exampleMonkeys()
			keepAway(monkeys, test.rounds, test.relief)

			var got []int
			passed := true
			for i := 0; i < len(monkeys); i++ {
				got = append(got, monkeys[i].inspections)
				if got[i] != test.want[i] {
					passed = false
				}
			}

			if !passed {
				t.Errorf("Expected:\n>>>%v<<<\nGot:\n>>>%v<<<", test.want, got)
			}
		})
	}
}

func exampleMonkeys() []monkey {
	return []monkey{
		{
			items: []int{79, 98},
			operation: func(old int) int {
				return old * 19
			},
			testMod:   23,
			passTrue:  2,
			passFalse: 3,
		},
		{
			items: []int{54, 65, 75, 74},
			operation: func(old int) int {
				return old + 6
			},
			testMod:   19,
			passTrue:  2,
			passFalse: 0,
		},
		{
			items: []int{79, 60, 97},
			operation: func(old int) int {
				return old * old
			},
			testMod:   13,
			passTrue:  1,
			passFalse: 3,
		},
		{
			items: []int{74},
			operation: func(old int) int {
				return old + 3
			},
			testMod:   17,
			passTrue:  0,
			passFalse: 1,
		},
	}
}
