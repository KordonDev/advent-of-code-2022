package main

import "fmt"

func main() {
	printMax("./1-example.txt")
	printMax("1.txt")

	printTop3Max("./1-example.txt")
	printTop3Max("1.txt")
}

func printMax(filename string) {
	lines := readFile(filename)
	max := getMaxValues(lines)

	fmt.Printf("Max for %s is %d.\n", filename, max[0])
}

func printTop3Max(filename string) {
	lines := readFile(filename)
	max := getMaxValues(lines)

	fmt.Printf("Max top 3 for %s is %d.\n", filename, max[0]+max[1]+max[2])
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
		currentMax = currentMax + stringToInt(value)
	}
	sortMaxToMin(&allMaxSorted)
	return allMaxSorted
}
