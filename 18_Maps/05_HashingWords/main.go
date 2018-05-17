package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// making a get request to a book resource
	res, err := http.Get("https://www.gutenberg.org/files/2701/2701-0.txt")
	// checking if the get request returned any error
	if err != nil {
		log.Fatal(err)
	}

	// reading the response body
	scanner := bufio.NewScanner(res.Body)
	// close the response body at the end of the main function
	defer res.Body.Close()
	// splitting the words in scanner
	scanner.Split(bufio.ScanWords)
	// slice to hold counts of len and capacity 200
	buckets := make([]int, 250)
	// loop over the words
	for scanner.Scan() {
		n := Hashbucket(scanner.Text())
		// incrementing the counter for the value of the character here
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
}

// Hashbucket : returns the ASCII value of the first letter of the word
func Hashbucket(word string) int {
	return int(word[0])
}
