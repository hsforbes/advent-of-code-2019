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

		// fmt.Println(line)

		if err == io.EOF {
			break
		}

		moduleMass, err := strconv.Atoi(line)

		fmt.Printf("\nOriginal mass:\t%d\t\t", moduleMass)

		fuelAddend := calculateFuel(moduleMass)
		fmt.Printf("\nAdded fuel: %d", fuelAddend)
		fuelSum += fuelAddend

	}
	fmt.Printf("\n\nCalculated fuel to carry modules: %d", fuelSum)
	// 4900529 too low
	// 4901534 too high

}

func calculateFuel(mass int) int {
	fuelReq := mass/3 - 2
	// fmt.Printf("%d", fuelReq)

	if fuelReq < 3 {
		if fuelReq < 1 {
			fmt.Printf("0")
			return 0
		} else {
			fmt.Printf("%d", fuelReq)
			return fuelReq
		}
	} else {
		fmt.Printf("%d + ", fuelReq)
		return fuelReq + calculateFuel(fuelReq)
	}
}

// func calculateFuelFuel(mass int) int {
// 	fuelReq := mass / 3 - 2

// 	if fuelReq < 3 {
// 		return fuelReq
// 	} else {
// 		return fuelReq +
// 	}
// }
