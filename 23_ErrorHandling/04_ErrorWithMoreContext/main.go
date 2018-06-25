package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-12.12)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("Negative square root number: %v", f)
	}

	return 12, nil
}
