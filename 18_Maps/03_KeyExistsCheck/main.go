package main

import "fmt"

func main() {
	myMapList := map[int]string{
		0: "luffy",
		1: "zoro",
		2: "sanji",
		3: "nami",
		4: "usopp",
	}

	fmt.Println(myMapList)

	if val, exists := myMapList[2]; exists {
		delete(myMapList, 2)
		fmt.Println("deleted: ", val)
	} else {
		fmt.Println("the value does not exist")
		fmt.Println("exists check: ", exists)
	}

	fmt.Println(myMapList)
}
