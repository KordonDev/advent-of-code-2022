package main

import (
	"fmt"
	"github.com/kordondev/advent-of-code-2022/helper"
	"strings"
)

func main() {
	startWindow4WithoutDuplicates("6-example.txt")
	startWindow4WithoutDuplicates("6.txt")

	startWindow14WithoutDuplicates("6-example.txt")
	startWindow14WithoutDuplicates("6.txt")
}

func startWindow14WithoutDuplicates(filename string) {
	lines := helper.ReadFile(filename)
	sequence := strings.Split(lines[0], "")
	for i := 13; i < len(sequence); i++ {
		var window []string
		for j := 0; j < 14; j++ {
			window = append(window, sequence[i-j])
		}
		if !hasDuplicates(window) {
			fmt.Printf("Sequence index is %d\n", i+1)
			break
		}
	}
}

func startWindow4WithoutDuplicates(filename string) {
	lines := helper.ReadFile(filename)
	sequence := strings.Split(lines[0], "")
	for i := 3; i < len(sequence); i++ {
		window := []string{sequence[i-3], sequence[i-2], sequence[i-1], sequence[i]}
		if !hasDuplicates(window) {
			fmt.Printf("Sequence index is %d\n", i+1)
			break
		}
	}
}

func hasDuplicates(s []string) bool {
	var m = make(map[string]bool)
	for _, v := range s {
		if m[v] {
			return true
		}
		m[v] = true
	}
	return false
}
