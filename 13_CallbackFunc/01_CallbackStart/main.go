package main

import "fmt"

func main() {
	// calling the visit function which takes in slice and callback function
	visit([]int{1, 2, 3, 4, 5, 6}, func(n int) {
		fmt.Println(n)
	})
}

// declaring the function with parameters for slice and callback function which takes an integer
func visit(numbers []int, callback func(int)) {
	// looping through numbers and calling the callback function
	for _, n := range numbers {
		callback(n)
	}
}
