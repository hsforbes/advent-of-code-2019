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
	// 3268951
}

func calculateFuel(unaccountedMass int) int {
	fuelReq := unaccountedMass / 3 - 2

	if fuelReq < 3 {
		return fuelReq
	} else {
		return fuelReq + calculateFuel(fuelReq)
	}
}