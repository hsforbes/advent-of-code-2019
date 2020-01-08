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
	findShortestIntersection(matchingCoordinates)
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
	x, y, distance int
}

type matchingCoordinate struct {
	x, y, distance1, distance2 int
}

func makeCoordinatesForAllSteps(wirePath []PathStep) []Coordinate {
	wireCoordinates := make([]Coordinate, 0)
	wireTipCoordinate := Coordinate{x: 0, y: 0, distance: 0}

	for i := 0; i < len(wirePath); i++ {
		var coordinatesForOneStep []Coordinate
		coordinatesForOneStep, wireTipCoordinate = makeCoordinatesForOneStep(wirePath[i], wireTipCoordinate)
		wireCoordinates = append(wireCoordinates, coordinatesForOneStep...)
	}

	// remove the first element. It's the 0,0 coordinate
	wireCoordinates = wireCoordinates[1:]

	return wireCoordinates
}

func makeCoordinatesForOneStep(pathStep PathStep, startCoordinate Coordinate) ([]Coordinate, Coordinate) {
	wireCoordinates := make([]Coordinate, 0)
	var newWireTipCoordinate = startCoordinate

	if strings.Compare(pathStep.direction, "U") == 0 {
		newWireTipCoordinate.y = startCoordinate.y + pathStep.magnitude
		for i := 0; i < pathStep.magnitude; i++ {
			newCoordinate := Coordinate{x: startCoordinate.x, y: startCoordinate.y + i, distance: startCoordinate.distance + i}
			wireCoordinates = append(wireCoordinates, newCoordinate)
		}
	} else if strings.Compare(pathStep.direction, "D") == 0 {
		newWireTipCoordinate.y = startCoordinate.y - pathStep.magnitude
		for i := 0; i < pathStep.magnitude; i++ {
			newCoordinate := Coordinate{x: startCoordinate.x, y: startCoordinate.y - i, distance: startCoordinate.distance + i}
			wireCoordinates = append(wireCoordinates, newCoordinate)
		}
	} else if strings.Compare(pathStep.direction, "L") == 0 {
		newWireTipCoordinate.x = startCoordinate.x - pathStep.magnitude
		for i := 0; i < pathStep.magnitude; i++ {
			newCoordinate := Coordinate{x: startCoordinate.x - i, y: startCoordinate.y, distance: startCoordinate.distance + i}
			wireCoordinates = append(wireCoordinates, newCoordinate)
		}
	} else if strings.Compare(pathStep.direction, "R") == 0 {
		newWireTipCoordinate.x = startCoordinate.x + pathStep.magnitude
		for i := 0; i < pathStep.magnitude; i++ {
			newCoordinate := Coordinate{x: startCoordinate.x + i, y: startCoordinate.y, distance: startCoordinate.distance + i}
			wireCoordinates = append(wireCoordinates, newCoordinate)
		}
	}

	return wireCoordinates, newWireTipCoordinate
}

func findMatchingCoordinates(wire1Coordinates []Coordinate, wire2Coordinates []Coordinate) []matchingCoordinate {
	matchingCoordinates := make([]matchingCoordinate, 0)

	for i := 0; i < len(wire1Coordinates); i++ {
		for j := 0; j < len(wire2Coordinates); j++ {
			if wire1Coordinates[i].x == wire2Coordinates[j].x {
				if wire1Coordinates[i].y == wire2Coordinates[j].y {
					intersection := matchingCoordinate{wire1Coordinates[i].x, wire1Coordinates[i].y, wire1Coordinates[i].distance, wire2Coordinates[j].distance}
					matchingCoordinates = append(matchingCoordinates, intersection)
					// matchingCoordinates = append(matchingCoordinates, wire1Coordinates[i])
				}
			}
		}
	}
	return matchingCoordinates
}

func findShortestIntersection(matchingCoordinates []matchingCoordinate) {

	// Set it to the max int and look for something smaller
	smallestDistance := math.MaxInt32

	for i := 0; i < len(matchingCoordinates); i++ {
		distance := calculateDistancesToIntersection(matchingCoordinates[i])
		if distance < smallestDistance {
			smallestDistance = distance
		}
	}

	fmt.Printf("\nFound smallest distance = %d", smallestDistance)
	// 87 is wrong

}

func calculateDistancesToIntersection(matchingCoordinate matchingCoordinate) int {
	return matchingCoordinate.distance1 + matchingCoordinate.distance2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readLine(reader *bufio.Reader) string {

	lineBytes, _, err := reader.ReadLine()
	line := string(lineBytes)

	if err == io.EOF {
		panic(err)
	}
	return line
}
