package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// making a get request to get the book of sherlock holmes
	res, err := http.Get("https://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	// check if any error is returned
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	// defer closing of the response body till the end of main
	defer res.Body.Close()
	// split in terms of words
	scanner.Split(bufio.ScanWords)
	// Creating a slice of slice of strings to hold slices of words
	buckets := make([][]string, 12)
	// preparing bucket to hold slice of words
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	// looping over the words
	for scanner.Scan() {
		word := scanner.Text()
		// calling the hash function to get the bucket value
		n := hashbucket(word, 12)
		// storing to the bucket value the word
		buckets[n] = append(buckets[n], word)
	}

	// printing the len of words in each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}

	// checking words present in one of the buckets
	fmt.Println(buckets[2])
}

func hashbucket(word string, buckets int) int {
	var sum int
	// looping through each character in the word and calculating the sum of each character
	for _, v := range word {
		sum += int(v)
	}
	// return remainder of the sum with the number of buckets for even distribution
	return sum % buckets
}
