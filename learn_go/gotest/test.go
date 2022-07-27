package main

import (
	"fmt"
	"strings"
)
//reverse a string

func main() {
	fmt.Printf(recombine(byter("yellow")))
}

func byter(x string) []string {

	output := make([]string, len(x))
	
	for pos, char := range x {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
		output = append(output, string(char))
	}

	return output
}

func recombine(x []string) string {
	
	var sbstr strings.Builder
	
	for i :=0; i < len(x); i++ {
		sbstr.WriteString(x[i])
	}

	return sbstr.String()
}