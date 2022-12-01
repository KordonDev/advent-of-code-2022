package main

import (
	"fmt"
	"github.com/kordondev/advent-of-code-2022/helper"
)

func main() {
	printMax("1-example.txt")
	printMax("1.txt")

	printTop3Max("1-example.txt")
	printTop3Max("1.txt")
}

func printMax(filename string) {
	lines := helper.ReadFile(filename)
	max := getMaxValues(lines)

	fmt.Printf("Max is %d.\n\n", max[0])
}

func printTop3Max(filename string) {
	lines := helper.ReadFile(filename)
	max := getMaxValues(lines)

	fmt.Printf("Max top 3 is %d.\n\n", max[0]+max[1]+max[2])
}

func getMaxValues(allValues []string) []int {

	var allMaxSorted []int
	currentMax := 0

	for _, value := range allValues {
		if value == "" {
			allMaxSorted = append(allMaxSorted, currentMax)
			currentMax = 0
			continue
		}
		currentMax = currentMax + helper.ToInt(value)
	}
	helper.SortMaxToMin(&allMaxSorted)
	return allMaxSorted
}
