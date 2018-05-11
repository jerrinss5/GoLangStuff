package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// calling a url
	res, error := http.Get("http://www.google.com")
	// if some issue in the get request error would pop and would be handled
	if error != nil {
		log.Fatal(error)
	}
	// reading the response body using ioutil
	page, error := ioutil.ReadAll(res.Body)
	// error handling for reading the response body
	if error != nil {
		log.Fatal(error)
	}

	res.Body.Close()
	fmt.Printf("%s", page)
}
