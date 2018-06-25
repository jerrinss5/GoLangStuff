package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-12)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("negative square root error")
	}

	return 12, nil
}
