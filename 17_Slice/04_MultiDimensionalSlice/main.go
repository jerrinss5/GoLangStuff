package main

import "fmt"

func main() {
	// storing a slice of slice of string in records
	records := make([][]string, 0)

	// student 1
	student1 := make([]string, 4)
	student1[0] = "Roronoa"
	student1[1] = "Zoro"
	student1[2] = "98.00"
	student1[3] = "95.00"

	// appending the student1 details to the records
	records = append(records, student1)

	// student 2
	student2 := make([]string, 4)
	student2[0] = "Monkey"
	student2[1] = "Luffy"
	student2[2] = "100.00"
	student2[3] = "98.00"

	// appending the student2 details to the records
	records = append(records, student2)

	fmt.Println(records)
}
