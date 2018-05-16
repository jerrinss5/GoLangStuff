package main

import "fmt"

func main() {
	var largestNumber = greatestNumber(3, 2, 5, 13, 6, 7, 11, 4)
	fmt.Println(largestNumber)
}

func greatestNumber(numbers ...int) int {
	var largeNumber int
	for _, n := range numbers {
		if n >= largeNumber {
			largeNumber = n
		}
	}

	return largeNumber
}
