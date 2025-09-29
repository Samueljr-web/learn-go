package main

import "fmt"

func main() {
	age := 18
	name := "Samuel"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line! \n")

	// Println
	fmt.Println("hello")
	fmt.Println("goodbye samuel!")
	fmt.Println("My age is", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name) // display the variable is default format
	fmt.Printf("my age is %q and my name is %q \n", age, name) // add quotes to string variables only
	fmt.Printf("age is of type %T \n", age) // checks the datatype of a value
	fmt.Printf("you score %f points! \n", 225.55) // adds floating point 
	fmt.Printf("you score %0.1f points! \n", 225.55) //round up 1
	fmt.Printf("you score %% points! \n")

	// Sprintf (save formatted strings)
	str := fmt.Sprintf("my age is %v and my name is %v \n", age, name) 
	println("the saved string is:", str)


}