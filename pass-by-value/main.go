package main

import "fmt"


func updateName(x string) {
	x = "yo"
}

func updateMenu(y map[string]float64) {
    y["cofee"] = 5000
}

func main() {
	// Group A types -> string, bools, floats, ints, arrays, struct
	name := "samuel"

    updateName(name)
	
	fmt.Println(name)

	//Group B types -> slices, maps, functions

	menu := map[string]float64 {
		"rice": 5499,
		"burger": 2550,
	}
   updateMenu(menu)
   fmt.Println(menu)
}
