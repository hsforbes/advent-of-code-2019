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
		
		line := readLine(reader)

		fmt.Printf("%s \n", line)
	}

}

func readLine(reader *bufio.Reader) string {

	lineBytes, _, err := reader.ReadLine()
	line := string(lineBytes)

	if err == io.EOF {
		panic(err)
	}
	return line
}
