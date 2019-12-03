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

	for {
		lineBytes, _, err := reader.ReadLine()
		line := string(lineBytes)

		if err == io.EOF {
			break
		}

		fmt.Printf("%s \n", line)

		inputStrings := strings.Split(line, ",")
		inputInts := makeIntInput(inputStrings)

		runIntcode(inputInts, 0)
		// 234699 too low
	}
}

func runIntcode(program []int, opcodePosition int) []int {
	opcode := program[opcodePosition]

	if opcode == 99 {
		fmt.Println("Found opcode 99. Stopping. Final program state:")
		fmt.Println(program)
		return program
	} else if opcode == 1 {
		program[program[opcodePosition+3]] = program[program[opcodePosition+1]] + program[program[opcodePosition+2]]
	} else if opcode == 2 {
		program[program[opcodePosition+3]] = program[program[opcodePosition+1]] * program[program[opcodePosition+2]]
	} else {
		fmt.Println("It a sploded")
	}

	return runIntcode(program, opcodePosition+4)
}

func makeIntInput(inputStrings []string) []int {
	inputInts := make([]int, len(inputStrings))
	var err error
	for i := 0; i < len(inputStrings); i++ {
		inputInts[i], err = strconv.Atoi(inputStrings[i])
	}
	if err != nil {
		panic(err)
	}
	return inputInts
}
