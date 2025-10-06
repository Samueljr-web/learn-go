package main

import (
	"fmt"
)

// func sayGreeting(n string){
// 	fmt.Printf("Good morning %v \n", n)
// }

// func sayGoodbye(n string){
// 	fmt.Printf("Goodbye %v \n", n)
// }

// func cycleNames(s []string, f func(string)){ //passing a function as params
// 	for _,v := range s{
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }

//my example

func sayGoodday(name string) {
	fmt.Printf("Good day %v \n", name)
}

func loopName(names []string, f func(string)){
   for _,value := range names{
	f(value)
   }
}


func main() {
//    sayGreeting("Samuel")
//    sayGoodbye("OX")
//    cycleNames([]string{"ox", "ad", "sam"}, sayGreeting)
//    cycleNames([]string{"ox", "ad", "sam"}, sayGoodbye)

//    a1 := circleArea(10.5)
//    a2 := circleArea(20.5)

//    fmt.Println(a1, a2) //without formatting
//    fmt.Printf("circle 1 is %0.3f circle 2 is %0.3f \n", a1, a2) //formatted to 3dp


   //my examples
   sayGoodday("sam")
   loopName([]string{"ad", "israel", "jacob", "marie"}, sayGoodday)
}

