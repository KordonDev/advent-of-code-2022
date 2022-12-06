package main

import (
	"fmt"
	"github.com/kordondev/advent-of-code-2022/helper"
	"strings"
)

func main() {
	calculateTopStackAfterMoving("5-example.txt")
	calculateTopStackAfterMoving("5.txt")

	calculateTopStackAfterMultiMoving("5-example.txt")
	calculateTopStackAfterMultiMoving("5.txt")
}

func calculateTopStackAfterMoving(filename string) {
	lines := helper.ReadFile(filename)

	betweenStackAndMovement := 0
	for i, l := range lines {
		if l == "" {
			betweenStackAndMovement = i
			break
		}
	}

	stacks := getStartStacks(lines[:betweenStackAndMovement])
	movements := getMovements(lines[betweenStackAndMovement+1:])

	stacks = moveStacks(stacks, movements)

	result := ""
	for _, s := range stacks {
		result = result + s[len(s)-1]
	}
	fmt.Printf("Resuld: %s from %v\n", result, stacks)
}

func calculateTopStackAfterMultiMoving(filename string) {
	lines := helper.ReadFile(filename)

	betweenStackAndMovement := 0
	for i, l := range lines {
		if l == "" {
			betweenStackAndMovement = i
			break
		}
	}

	stacks := getStartStacks(lines[:betweenStackAndMovement])
	movements := getMovements(lines[betweenStackAndMovement+1:])

	stacks = moveMultiStacks(stacks, movements)

	result := ""
	for _, s := range stacks {
		result = result + s[len(s)-1]
	}
	fmt.Printf("Resuld: %s from %v\n", result, stacks)
}

func getStartStacks(lines []string) [][]string {
	var stacks [][]string
	for i := len(lines) - 1; i >= 0; i-- {
		columns := strings.Split(lines[i], "")
		for j := 0; j < len(columns); j++ {
			readIndex := j + 1 + (j * 3)
			if readIndex >= len(columns) {
				break
			}

			if i == len(lines)-1 {
				stacks = append(stacks, []string{})
			} else if columns[readIndex] != " " {
				stacks[j] = append(stacks[j], columns[readIndex])
			}
		}
	}
	return stacks
}

type movement struct {
	from   int
	to     int
	amount int
}

func getMovements(lines []string) []movement {
	var movements []movement
	for _, l := range lines {
		c := strings.Split(l, " ")
		movements = append(movements, movement{
			amount: helper.ToInt(c[1]),
			from:   helper.ToInt(c[3]) - 1,
			to:     helper.ToInt(c[5]) - 1,
		})
	}
	return movements
}

func moveStacks(stacks [][]string, movements []movement) [][]string {
	for _, m := range movements {
		for i := 0; i < m.amount; i++ {
			stacks[m.to] = append(stacks[m.to], stacks[m.from][len(stacks[m.from])-1])
			stacks[m.from] = stacks[m.from][:len(stacks[m.from])-1]
		}

	}
	return stacks
}

func moveMultiStacks(stacks [][]string, movements []movement) [][]string {
	for _, m := range movements {
		move := stacks[m.from][len(stacks[m.from])-m.amount:]
		stacks[m.to] = append(stacks[m.to], move...)
		stacks[m.from] = stacks[m.from][:len(stacks[m.from])-m.amount]

	}
	return stacks
}
