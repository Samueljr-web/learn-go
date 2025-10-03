package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string){
	fmt.Printf("Good morning %v \n", n)
}

func sayGoodbye(n string){
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(s []string, f func(string)){ //passing a function as params
	for _,v := range s{
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
//    sayGreeting("Samuel")
//    sayGoodbye("OX")
//    cycleNames([]string{"ox", "ad", "sam"}, sayGreeting)
//    cycleNames([]string{"ox", "ad", "sam"}, sayGoodbye)

   a1 := circleArea(10.5)
   a2 := circleArea(20.5)

   fmt.Println(a1, a2) //without formatting
   fmt.Printf("circle 1 is %0.3f circle 2 is %0.3f \n", a1, a2) //formatted to 3dp
}

