package main

import (
	"testing"

	"go-aoc/internal/test"
)

func TestPartOne(t *testing.T) {
	tests := []test.ChallengeTest{
		{
			Name:     "(())",
			Input:    `(())`,
			Expected: 0,
		},
		{
			Name:     "(((",
			Input:    `(((`,
			Expected: 3,
		},
		{
			Name:     "(()(()(",
			Input:    `(()(()(`,
			Expected: 3,
		},
		{
			Name:     "())",
			Input:    `())`,
			Expected: -1,
		},
		{
			Name:     ")())())",
			Input:    `)())())`,
			Expected: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := partOne(tt.Input)
			if got != tt.Expected {
				t.Errorf("part1() = %d; expected %d", got, tt.Expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []test.ChallengeTest{
		{
			Name:     ")",
			Input:    `)`,
			Expected: 1,
		},
		{
			Name:     "()())",
			Input:    `()())`,
			Expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := partTwo(tt.Input)
			if got != tt.Expected {
				t.Errorf("part2() = %d; expected %d", got, tt.Expected)
			}
		})
	}
}
