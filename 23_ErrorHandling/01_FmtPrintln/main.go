package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")

	if err != nil {
		fmt.Println("Error Occurred ", err)
		// log.Println("Error Occurred: ", err)
		// log.Fatalln(err)
		// panic(err)
	}
}
