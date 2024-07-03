package main

import (
	"fmt"
	"strings"
)

// Check S shots
func checkBehavior(S string) string {
	if len(S) == 0 {
		return ""
	}
	S = strings.ToUpper(S)
	// Check First Chareter R
	if strings.HasPrefix(S, "R") {
		return "Bad boy"
	}

	shots := 0
	// For Loop Calculate Shots
	for _, shot := range S {
		if shot == 'S' {
			shots++
		} else if shot == 'R' && shots > 0 {
			shots--
		}
	}
	if shots > 0 {
		return "Bad boy"
	}
	return "Good boy"
}

func main() {
	tests := []string{
		"SRSSRRR",
		"RSSRR",
		"SSSRRRRS",
		"SRRSSR",
		"SSRSRRR",
	}

	for _, test := range tests {
		result := checkBehavior(test)
		fmt.Printf("INPUT %s ==> Output %s\n", test, result)
	}
}
