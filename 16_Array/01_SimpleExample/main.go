package main

import "fmt"

func main() {
	// declaring an array of strings of size 58
	var alphabets [58]string

	// printing the output before assignment
	fmt.Println(alphabets)
	fmt.Println(len(alphabets))
	fmt.Println(alphabets[42])

	// assigning the asci value to the array of string
	for i := 65; i <= 122; i++ {
		alphabets[i-65] = string(i)
	}

	// printing the output after assignment
	fmt.Println(alphabets)
	fmt.Println(len(alphabets))
	fmt.Println(alphabets[42])
}
