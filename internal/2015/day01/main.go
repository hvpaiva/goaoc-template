package main

import (
	_ "embed"
	"go-aoc/pkg/parser"
	"log"

	"github.com/hvpaiva/goaoc"
)

//go:embed input.txt
var input string

func main() {
	err := goaoc.Run(parser.Normalize(input), partOne, partTwo)
	if err != nil {
		log.Fatalf("Error running AOC: %v", err)
	}
}

func partOne(input string) int {
	var lvl int
	for _, str := range input {
		lvl += normalize(str)
	}

	return lvl
}

func partTwo(input string) int {
	var lvl int

	for i, str := range input {
		lvl += normalize(str)

		if lvl < 0 {
			return i + 1
		}
	}

	return 0
}

func normalize(n rune) int {
	if n == '(' {
		return 1
	}

	return -1
}
