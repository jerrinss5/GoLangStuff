package main

import (
	"errors"
	"log"
)

var errSqrtNegative = errors.New("Negative square root error")

func main() {
	_, err := sqrt(-12)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errSqrtNegative
	}

	return 12, nil
}
