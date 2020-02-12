package tdd

import (
	"fmt"
)

/*
Multiples gets multiples of 3 and 5 greater than or equal to start and less than or equal to end`
*/
func Multiples(start, end int) []int {
	result := make([]int, 0)
	for i := start; i <= end; i++ {
		if i%3 == 0 || i%5 == 0 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	fmt.Printf("anystring")
}
