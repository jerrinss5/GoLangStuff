package main

import "fmt"

func main() {
	var sum int
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	fmt.Println("Sum of multiples of 3 or 5 under 1000 is: ", sum)
}
