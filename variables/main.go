package main

import "fmt"

var outside string = "outside"
func main() {
	// fmt.Println("Hello, Samuel")

	var nameOne string = "samuel"
	var nameTwo = "Ad"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "ayo"
	nameThree = "odun"

	nameFour := "samuel" //cant be used outside of a function
 
	fmt.Println(nameFour)

	//ints
	var ageOne int = 18
	var ageTwo = 20
	ageThree := 30

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint16 = 256

	//floats 
	var scoreOne float32 = 25.9
	var scoreTwo float64 = 39999848
	scoreThree := 4040094

    fmt.Println(scoreOne,scoreTwo,scoreThree, outside)
}