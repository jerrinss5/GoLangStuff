package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// calling a url
	res, _ := http.Get("http://www.google.com")
	// reading the response body using ioutil
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
