package main

import "fmt"



func main() {
	menu := map[string]float64 {
		"rice": 5000.00,
		"spag": 2500.00,
		"yam": 3550.00,
	}
	fmt.Println(menu)
	fmt.Println(menu["yam"])

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