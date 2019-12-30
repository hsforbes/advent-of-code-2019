package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	wire1Coordinates := makeCoordinatesForAllSteps(wire1Path)
	wire2Coordinates := makeCoordinatesForAllSteps(wire2Path)

	matchingCoordinates := findMatchingCoordinates(wire1Coordinates, wire2Coordinates)
	findClosestIntersection(matchingCoordinates)
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

type Coordinate struct {
	x, y int
}

func makeCoordinatesForAllSteps(wirePath []PathStep) []Coordinate {
	wireTipCoordinate := Coordinate{x: 0, y: 0}
	wireCoordinates := make([]Coordinate, 0)
	for i := 0; i < len(wirePath); i++ {
		var coordinatesForOneStep []Coordinate
		coordinatesForOneStep, wireTipCoordinate = makeCoordinatesForOneStep(wirePath[i], wireTipCoordinate)
		wireCoordinates = append(wireCoordinates, coordinatesForOneStep...)
	}

	return wireCoordinates
}

func makeCoordinatesForOneStep(pathStep PathStep, startCoordinate Coordinate) ([]Coordinate, Coordinate) {
	wireCoordinates := make([]Coordinate, 0)
	var newWireTipCoordinate = startCoordinate

	if strings.Compare(pathStep.direction, "U") == 0 {
		newWireTipCoordinate.y = startCoordinate.y + pathStep.magnitude
		for i := 0; i < pathStep.magnitude; i++ {
			newCoordinate := Coordinate{x: startCoordinate.x, y: startCoordinate.y + i}
			wireCoordinates = append(wireCoordinates, newCoordinate)
		}
	} else if strings.Compare(pathStep.direction, "D") == 0 {
		newWireTipCoordinate.y = startCoordinate.y - pathStep.magnitude
		for i := 0; i < pathStep.magnitude; i++ {
			newCoordinate := Coordinate{x: startCoordinate.x, y: startCoordinate.y - i}
			wireCoordinates = append(wireCoordinates, newCoordinate)
		}
	} else if strings.Compare(pathStep.direction, "L") == 0 {
		newWireTipCoordinate.x = startCoordinate.x - pathStep.magnitude
		for i := 0; i < pathStep.magnitude; i++ {
			newCoordinate := Coordinate{x: startCoordinate.x - i, y: startCoordinate.y}
			wireCoordinates = append(wireCoordinates, newCoordinate)
		}
	} else if strings.Compare(pathStep.direction, "R") == 0 {
		newWireTipCoordinate.x = startCoordinate.x + pathStep.magnitude
		for i := 0; i < pathStep.magnitude; i++ {
			newCoordinate := Coordinate{x: startCoordinate.x + i, y: startCoordinate.y}
			wireCoordinates = append(wireCoordinates, newCoordinate)
		}
	}

	return wireCoordinates, newWireTipCoordinate
}

// TODO Make this efficient
func findMatchingCoordinates(wire1Coordinates []Coordinate, wire2Coordinates []Coordinate) []Coordinate {
	matchingCoordinates := make([]Coordinate, 0)

	for i := 0; i < len(wire1Coordinates); i++ {
		for j := 0; j < len(wire2Coordinates); j++ {
			if wire1Coordinates[i].x == wire2Coordinates[j].x {
				if wire1Coordinates[i].y == wire2Coordinates[j].y {
					matchingCoordinates = append(matchingCoordinates, wire1Coordinates[i])
				}
			}
		}
	}
	return matchingCoordinates
}

func findClosestIntersection(matchingCoordinates []Coordinate) {

	smallestDistance := math.MaxInt32
	// smallestDistanceCoordinate := Coordinate{x:0, y:0}

	for i := 0; i < len(matchingCoordinates); i++ {
		distance := matchingCoordinates[i].x + matchingCoordinates[i].y
		if distance < smallestDistance {
			smallestDistance = distance
			// smallestDistanceCoordinate = matchingCoordinates[i]
		}
	}

	fmt.Printf("\nFound smallest distance = %d", smallestDistance)
}

func readLine(reader *bufio.Reader) string {

	lineBytes, _, err := reader.ReadLine()
	line := string(lineBytes)

	if err == io.EOF {
		panic(err)
	}
	return line
}
