package main

import (
	"fmt"
	"github.com/kordondev/advent-of-code-2022/helper"
	"strings"
)

type Space struct {
	start int
	end   int
}

func main() {
	calculateIncludes("4-example.txt")
	calculateIncludes("4.txt")

	calculateOverlap("4-example.txt")
	calculateOverlap("4.txt")
}

func calculateOverlap(filename string) int {
	lines := helper.ReadFile(filename)

	overlap := 0
	spaces := parseSpaces(lines)
	for _, s := range spaces {
		if isOverlap(s) {
			overlap = overlap + 1
		}
	}
	fmt.Printf("Overlaps %d\n", overlap)
	return overlap
}

func isOverlap(s []Space) bool {
	if s[0].start <= s[1].start && s[0].end >= s[1].end {
		return true
	}
	if s[1].start <= s[0].start && s[1].end >= s[0].end {
		return true
	}
	if s[0].start <= s[1].start && s[1].start <= s[0].end {
		return true
	}
	if s[0].start <= s[1].end && s[1].end <= s[0].end {
		return true
	}
	return false
}

func calculateIncludes(filename string) {
	lines := helper.ReadFile(filename)

	includes := 0
	spaces := parseSpaces(lines)
	for _, s := range spaces {
		if s[0].start <= s[1].start && s[0].end >= s[1].end {
			includes = includes + 1
			continue
		}
		if s[1].start <= s[0].start && s[1].end >= s[0].end {
			includes = includes + 1
		}
	}
	fmt.Printf("includes %d\n", includes)
}

func parseSpaces(lines []string) [][]Space {
	var result [][]Space
	for _, l := range lines {
		var s []Space
		spaces := strings.Split(l, ",")
		space1s := helper.ToInts(strings.Split(spaces[0], "-"))
		space2s := helper.ToInts(strings.Split(spaces[1], "-"))

		space1 := Space{
			start: space1s[0],
			end:   space1s[1],
		}
		s = append(s, space1)
		space2 := Space{
			start: space2s[0],
			end:   space2s[1],
		}
		s = append(s, space2)

		result = append(result, s)
	}
	return result
}
