package main

import (
	"fmt"
	"sort"
)

type people []string

// Len : to get the length of the slice
func (ps people) Len() int {
	return len(ps)
}

// Less : to check if the first index is less than the second
func (ps people) Less(i, j int) bool {
	return ps[i] < ps[j]
}

// Swap : to swap the values of the indexes using temp variable
func (ps people) Swap(i, j int) {
	var temp = ps[i]
	ps[i] = ps[j]
	ps[j] = temp
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)

	sort.Sort(studyGroup)

	fmt.Println(studyGroup)
}
