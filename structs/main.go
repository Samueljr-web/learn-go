package main

import "fmt"


func main() {
	type User struct {
		id int
		name string
		age int
	}

	u := User{name: "Sam", age: 18}
	fmt.Println(u.name) 
}
