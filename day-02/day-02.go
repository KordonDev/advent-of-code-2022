package main

import (
	"fmt"
	"github.com/kordondev/advent-of-code-2022/helper"
	"strings"
)

func main() {
	linesExample := helper.ReadFile("2-example.txt")
	calculatePoints(linesExample)

	lines := helper.ReadFile("2.txt")
	calculatePoints(lines)

	calculatePoints2(linesExample)
	calculatePoints2(lines)
}
func calculatePoints(rounds []string) {
	totalPoints := 0
	for _, line := range rounds {
		signs := strings.Split(line, " ")
		if signs[0] == "A" {
			opp := A{}
			totalPoints = totalPoints + opp.getPoints(signs[1])
		}
		if signs[0] == "B" {
			opp := B{}
			totalPoints = totalPoints + opp.getPoints(signs[1])
		}
		if signs[0] == "C" {
			opp := C{}
			totalPoints = totalPoints + opp.getPoints(signs[1])
		}
	}
	fmt.Printf("total points are %d\n", totalPoints)
}
func calculatePoints2(rounds []string) {
	totalPoints := 0
	for _, line := range rounds {
		signs := strings.Split(line, " ")
		if signs[0] == "A" {
			opp := A{}
			totalPoints = totalPoints + opp.getPoints2(signs[1])
		}
		if signs[0] == "B" {
			opp := B{}
			totalPoints = totalPoints + opp.getPoints2(signs[1])
		}
		if signs[0] == "C" {
			opp := C{}
			totalPoints = totalPoints + opp.getPoints2(signs[1])
		}
	}
	fmt.Printf("2: total points are %d\n", totalPoints)
}

// Rock
type A struct{}

// Paper
type B struct{}

// Sissour
type C struct{}

func (a A) getPoints(mySign string) int {
	if mySign == "X" {
		return 1 + 3
	}
	if mySign == "Y" {
		return 2 + 6
	}
	return 3 + 0
}

func (b B) getPoints(mySign string) int {
	if mySign == "X" {
		return 1 + 0
	}
	if mySign == "Y" {
		return 2 + 3
	}
	return 3 + 6
}
func (c C) getPoints(mySign string) int {
	if mySign == "X" {
		return 1 + 6
	}
	if mySign == "Y" {
		return 2 + 0
	}
	return 3 + 3
}

// x = lose, y = draw, z = win
func (a A) getPoints2(mySign string) int {
	if mySign == "X" {
		return 3 + 0
	}
	if mySign == "Y" {
		return 1 + 3
	}
	return 2 + 6
}

func (b B) getPoints2(mySign string) int {
	if mySign == "X" {
		return 1 + 0
	}
	if mySign == "Y" {
		return 2 + 3
	}
	return 3 + 6
}
func (c C) getPoints2(mySign string) int {
	if mySign == "X" {
		return 2 + 0
	}
	if mySign == "Y" {
		return 3 + 3
	}
	return 1 + 6
}
