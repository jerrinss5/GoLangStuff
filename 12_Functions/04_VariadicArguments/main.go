package main

import "fmt"

func main() {
	// converting the individual data elements to slice of float64
	data := []float64{43, 56, 87, 12, 45, 57}
	// sending individual elements of float64 from slice to variadic function
	n := average(data...)
	fmt.Println(n)
}

func average(values ...float64) float64 {
	var total float64

	for _, v := range values {
		total += v
	}

	return total / float64(len(values))
}
