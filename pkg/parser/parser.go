package parser

import (
	"log"
	"strings"
)

func Normalize(input string) string {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		log.Fatalf("input.txt is empty")
	}

	return input
}
