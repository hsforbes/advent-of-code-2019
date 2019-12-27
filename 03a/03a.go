package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Entry point
func main() {
	fmt.Println("Are you ready to go?")

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	line1 := readLine(reader)
	line2 := readLine(reader)

	wire1Path := makeWirePath(line1)
	wire2Path := makeWirePath(line2)

	maxDimensions := findMaxDimensions(wire1Path, wire2Path)
	gridWithIntersections := makeGridWithIntersections(maxDimensions, wire1Path, wire2Path)

	findClosestIntersection(gridWithIntersections)
}

type PathStep struct {
	direction string
	magnitude int
}

func makePathStep(step string) PathStep {
	direction := string(step[0])
	magnitude, err := strconv.Atoi(string(step[1:]))
	if err != nil {
		panic(err)
	}

	return PathStep{direction, magnitude}
}

func makeWirePath(line string) []PathStep {

	steps := strings.Split(line, ",")
	wirePath := make([]PathStep, len(steps))
	for i := 0; i < len(steps); i++ {
		wirePath[i] = makePathStep(steps[i])
	}

	return wirePath
}

// func makeGrid(wirePath []PathStep) [][]int {
// 	// TODO

// 	// Maybe try generating the grid in 4 quadrants, then combining (or not)
// 	// OR
// 	// Maybe calculate the maximum dimensions first
// 	return [][]int{}
// }

func makeGridWithIntersections(maxDimensions []int, wire1Path []PathStep, wire2Path []PathStep) [][]int {

	// TODO
	return [][]int{}
}

func findMaxDimensions(wire1Path []PathStep, wire2Path []PathStep) []int {
	// Up, down, left, right
	maxDimensions := []int{0, 0, 0, 0}

	// x, y
	coordinate := []int{0, 0}

	for i := 0; i < len(wire1Path); i++ {
		if strings.Compare(wire1Path[i].direction, "U") == 0 {
			coordinate[]

		} else if strings.Compare(wire1Path[i].direction, "D") == 0 {

		} else if strings.Compare(wire1Path[i].direction, "L") == 0 {

		} else if strings.Compare(wire1Path[i].direction, "R") == 0 {

		}
	}

	return maxDimensions
}

func findClosestIntersection(gridWithIntersections [][]int) {
	// fmt.Println("The end")
}

func readLine(reader *bufio.Reader) string {

	lineBytes, _, err := reader.ReadLine()
	line := string(lineBytes)

	if err == io.EOF {
		panic(err)
	}
	return line
}
