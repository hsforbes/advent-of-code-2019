package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Entry point
func main() {
	fmt.Println("Are you ready to go?")

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	fuelSum := 0

	for {
		lineBytes, _, err := reader.ReadLine()
		line := string(lineBytes)

		if err == io.EOF {
			break
		}

		moduleMass, err := strconv.Atoi(line)

		fuelSum += calculateFuel(moduleMass)

	}
	fmt.Printf("Calculated fuel to carry modules: %d", fuelSum)

}

func calculateFuel(moduleMass int) int {

	return moduleMass / 3 - 2
}