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
	line := readLine(reader)

	// fmt.Printf("%s \n", line)

	// inputStrings := strings.Split(line, ",")
	wire1Path := makeWirePath(line)
	wire1Grid := makeGrid(wire1Path)

	// Read Wire 2
	line = readLine(reader)
	wire2Path := makeWirePath(line)
	wire2Grid := makeGrid(wire2Path)

	gridWithIntersections := makeGridWithIntersections(wire1Grid, wire2Grid)

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

func makeGrid(wirePath []PathStep) [][]int {
	// TODO

	// Maybe try generating the grid in 4 quadrants, then combining (or not)
	return [][]int{}
}

func makeGridWithIntersections(wire1Grid [][]int, wire2Grid [][]int) [][]int {
	// TODO
	return [][]int{}
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
