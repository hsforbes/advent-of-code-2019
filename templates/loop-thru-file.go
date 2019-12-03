package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	}

}
