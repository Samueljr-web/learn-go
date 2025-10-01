package main

import (
	"fmt"
	"sort"
)

func main(){
	// greeting := "Hello goodday!"

	// //note: these methods returns a new string it doesnt alter the old one
	// fmt.Println(strings.Contains(greeting, "gooday")) //the "Contains" string method returns true or false
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hey")) 
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "dd"))
	// fmt.Println(strings.Split(greeting, " ")) //turns it to slices

	ages := []int{20, 10, 15, 30, 55 ,60 ,25}
	sort.Ints(ages) // maddd - if it was js(i laugh) 
	fmt.Println(ages)

	index := sort.SearchInts(ages, 55)
	fmt.Println(index)

	names := []string{"samuel", "ad", "ayo", "Aydee"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "ad"))
}
