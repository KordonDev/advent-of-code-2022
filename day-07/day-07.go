package main

import (
	"fmt"
	"github.com/kordondev/advent-of-code-2022/helper"
	"math"
	"strings"
)

func main() {
	calculateResult1("07-example.txt")
	calculateResult1("07.txt")

	findDirectoryToFreeUp("07-example.txt")
	findDirectoryToFreeUp("07.txt")
}

func findDirectoryToFreeUp(filename string) {
	lines := helper.ReadFile(filename)
	root := createFileSystem(lines)
	calculateSize(root)

	needToFreeUp := -70000000 + 30000000 + root.cumulatedSize

	fmt.Printf("Free up: %d\n", smallestDirectoryToFreeUp(root, needToFreeUp, math.MaxInt))

}

func smallestDirectoryToFreeUp(dir *Directory, minSize, currentSize int) int {
	smallestSize := currentSize
	for _, child := range dir.children {
		smallestSize = smallestDirectoryToFreeUp(child, minSize, smallestSize)
	}
	if dir.cumulatedSize >= minSize && dir.cumulatedSize < smallestSize {
		smallestSize = dir.cumulatedSize
	}
	return smallestSize
}

func calculateResult1(filename string) {
	lines := helper.ReadFile(filename)
	root := createFileSystem(lines)
	calculateSize(root)
	fmt.Printf("sum of small sizes: %d\n", sumOfSmallSizes(root))
}

func createFileSystem(commands []string) *Directory {
	var root *Directory
	var currentDirectory *Directory
	for _, c := range commands {
		if strings.HasPrefix(c, "$") {
			parts := strings.Split(c, " ")

			if c == "$ cd .." {
				currentDirectory = currentDirectory.parent
			} else if parts[1] == "cd" {
				_, dir := createDirectoryIfNecessary(currentDirectory, parts[2])
				currentDirectory = dir
				if root == nil {
					root = currentDirectory
				}
			}
		} else {
			if strings.HasPrefix(c, "dir") {
				parts := strings.Split(c, " ")
				created, dir := createDirectoryIfNecessary(currentDirectory, parts[1])
				if created {
					currentDirectory.children = append(currentDirectory.children, dir)
				}
			} else {
				parts := strings.Split(c, " ")
				file := File{
					name: parts[1],
					size: helper.ToInt(parts[0]),
				}
				currentDirectory.files = append(currentDirectory.files, &file)
			}
		}
	}
	return root
}

func sumOfSmallSizes(directory *Directory) int {
	size := 0
	for _, child := range directory.children {
		size = size + sumOfSmallSizes(child)
	}
	if directory.cumulatedSize <= 100000 {
		size = size + directory.cumulatedSize
	}
	return size
}

func calculateSize(directory *Directory) int {
	size := 0
	for _, child := range directory.children {
		size = size + calculateSize(child)
	}
	for _, file := range directory.files {
		size = size + file.size
	}
	directory.cumulatedSize = size
	return size
}

func createDirectoryIfNecessary(currentDir *Directory, dirName string) (bool, *Directory) {
	if currentDir != nil {
		for _, d := range currentDir.children {
			if d.name == dirName {
				return false, d
			}
		}
	}
	return true, &Directory{
		name:     dirName,
		parent:   currentDir,
		files:    make([]*File, 0),
		children: make([]*Directory, 0),
	}
}

type File struct {
	name string
	size int
}

type Directory struct {
	name          string
	parent        *Directory
	cumulatedSize int
	files         []*File
	children      []*Directory
}
