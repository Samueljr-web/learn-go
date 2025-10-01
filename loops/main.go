package main

import "fmt"

func main() {
	//basic loop
	x := 0
	for x < 5 {
		fmt.Println("value is", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value i:", i)
	}

	var names = []string{"sam", "ad", "ayo", "israel"}

   //looping through a string

	for i := 0; i < len(names); i++ {
      fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value )
	}

	for  _, value := range names {
		fmt.Printf("the value is %v \n", value)
	}
}
