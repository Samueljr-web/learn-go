package main

import "fmt"


func updateName(x *string) {
	*x = "yo"

	// fmt.Println("the memory location of x is:", &x)
}



func main() {
	name := "samuel"

    // updateName(name)

	// fmt.Println("the memory location of name is:", &name)
	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)


	updateName(m)
	fmt.Println(name)
}