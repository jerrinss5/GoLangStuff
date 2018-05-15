package main

import "fmt"

func main() {
	returnedSliceElements := visit([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
		return n > 1
	})
	fmt.Println(returnedSliceElements)
}

// here visit function in itself is returning a slice
// callback function is returning a boolean
func visit(numbers []int, callback func(int) bool) []int {
	sliceElements := []int{}
	for _, n := range numbers {
		if callback(n) {
			sliceElements = append(sliceElements, n)
		}
	}
	return sliceElements
}
