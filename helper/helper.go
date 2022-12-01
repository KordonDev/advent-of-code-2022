package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func ReadFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var texts []string
	for scanner.Scan() {
		texts = append(texts, scanner.Text())
	}
	file.Close()
	fmt.Printf("Read file %s\n", filePath)
	return texts
}

func ToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func ToInts(s []string) []int {
	var numbers []int
	for _, line := range s {
		numbers = append(numbers, ToInt(line))
	}
	return numbers
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// sort max to min
func SortMaxToMin(ints *[]int) {
	sort.Ints(*ints)
	sort.SliceStable(*ints, func(i, j int) bool {
		return j < i
	})
}

func SortMinToMax(ints *[]int) {
	sort.Ints(*ints)
}

func Slice(array []string, start int, end int, insert string) []string {
	a1 := append(array[:start], insert)
	if end >= len(array) {
		return a1
	}
	return append(a1, array[end:]...)
}

func SliceArray(array []string, start int, end int, insert []string) []string {
	a1 := append(array[:start], insert...)
	if end >= len(array) {
		return a1
	}
	return append(a1, array[end:]...)
}

func RemoveItems(list1 *[]string, list2 *[]string) ([]string, bool) {
	var result []string
	change := false
	for _, item1 := range *list1 {
		contains := false
		for _, item2 := range *list2 {
			if item1 == item2 {
				contains = true
				break
			}
		}
		if !contains {
			result = append(result, item1)
		} else {
			change = true
		}
	}
	return result, change
}

func SplitBy(list []string, divider string) [][]string {
	var result [][]string
	var currentList []string
	for _, v := range list {
		if v == divider {
			result = append(result, currentList)
			currentList = []string{}
			continue
		}
		currentList = append(currentList, v)
	}

	return result
}
