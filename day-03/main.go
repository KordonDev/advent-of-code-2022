package main

import (
	"fmt"
	"github.com/kordondev/advent-of-code-2022/helper"
	"strings"
)

func main() {
	calculatePriority("3-example.txt")
	calculatePriority("3.txt")

	priorityByElveGroups("3-example.txt")
	priorityByElveGroups("3.txt")
}

func priorityByElveGroups(filename string) {
	lines := helper.ReadFile(filename)
	var sameChar []string
	elveNumber := 1
	sameCharInGroup := make(map[string]int)
	for _, l := range lines {
		added := make(map[string]bool)
		lineList := strings.Split(l, "")
		for _, c := range lineList {
			if !added[c] {
				sameCharInGroup[c] = sameCharInGroup[c] + 1
				added[c] = true
			}
		}

		if elveNumber == 3 {
			for k := range sameCharInGroup {
				if sameCharInGroup[k] == 3 {
					sameChar = append(sameChar, k)
				}
			}
			sameCharInGroup = make(map[string]int)
			added = make(map[string]bool)
			elveNumber = 1
			continue
		}
		elveNumber = elveNumber + 1
	}

	result := 0
	for _, s := range sameChar {
		result = result + toPriority(s)
	}

	fmt.Printf("%d\n", result)
}

func calculatePriority(filename string) {
	lines := helper.ReadFile(filename)

	var sameChar []string
	for _, l := range lines {
		half := len(l) / 2
		firstHalf := make(map[string]bool)
		lineList := strings.Split(l, "")
		for _, c := range lineList[:half] {
			firstHalf[c] = true
		}
		for _, c := range lineList[half:] {
			if firstHalf[c] {
				sameChar = append(sameChar, c)
				break
			}
		}
	}

	result := 0
	for _, s := range sameChar {
		result = result + toPriority(s)
	}

	fmt.Printf("%d\n", result)

}

func toPriority(s string) int {
	ascii := s[0]
	if ascii < 95 {
		return int(ascii - "A"[0] + 27)
	}
	return int(ascii - "a"[0] + 1)
}
