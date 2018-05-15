package main

import "fmt"

func main() {
	// calling the visit function with slice value and funcition definition of int return type
	visit([]int{1, 2, 3, 4, 5, 6}, func(n int) int {
		return n
	})
}

// declaring visit function which takes in slice and function value
func visit(numbers []int, callback func(int) int) {
	for _, n := range numbers {
		fmt.Println(callback(n))
	}
}

// it is called callback because it is calling the function which passes the agrument
