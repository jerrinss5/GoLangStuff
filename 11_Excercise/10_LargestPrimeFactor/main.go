package main

import "fmt"
import "math"

func main() {
	// printing out the largest number
	fmt.Println(largestPrimeFactor(13195))
}

func largestPrimeFactor(n float64) float64 {
	var primeFactors []float64
	// looping through the numbers from 2 to n
	for i := 2.0; i <= n; i++ {
		// using math mod library to do modulus operation on the float numbers
		if math.Mod(n, i) == 0 {
			// reducing the number range
			n = n / i
			// appending the factor to the slice
			primeFactors = append(primeFactors, i)
			i--
		}
	}
	var largestPrimeFactorNumber float64
	// loopign through the prime factors to find the largest of the numbers
	for _, n := range primeFactors {
		if largestPrimeFactorNumber <= n {
			largestPrimeFactorNumber = n
		}
	}
	return largestPrimeFactorNumber
}

/*

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/
