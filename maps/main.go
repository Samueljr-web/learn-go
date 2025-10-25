package main

import "fmt"



func main() {
	menu := map[string]int {
		"a": 1,
		"b": 2,
	}
	// looping maps
	for i,v := range menu {
		fmt.Printf("%v is %v \n", i, v)
	}

	// ints as key type

	num := map[int]string {
		1: "a",
		2: "b",
		3: "c",
	}

	fmt.Println(num)
	fmt.Println(num[1])

	num[1] = "z"
	fmt.Println(num)
}