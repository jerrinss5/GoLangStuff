package main

import "fmt"

func main() {
	//average is the function which can take in any number of arguments
	averageValue := average(43, 56, 87, 12, 14, 15)
	fmt.Println(averageValue)
}

// array value of numbers are accepted as parameters
func average(values ...float64) float64 {
	var total float64

	for _, v := range values {
		total += v
	}

	return total / float64(len(values))
}
