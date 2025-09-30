package main

import "fmt"

func main() {
	//Arrays

	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}
    
	names := [4]string{"Samuel", "Ad", "Ay","Ox" }
	names[1] = "Adeyemi"
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))


	// slices (use arrays under the hood)

	var scores = []int{100, 50, 60}
	scores[2] = 40
	scores = append(scores, 585) // can only append a slice

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[1:]
	rangeThree := names[:3]
	fmt.Println(rangeOne,rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "test") // can also append range
	fmt.Println(rangeOne)



	//my example

	//array
	classMembers := [3]string{"samuel", "jacob", "onyi"}
	fmt.Println(classMembers)

	//slices
	classTutors := []string{"mr","mrs","dr"}
	fmt.Print(classTutors)
	classTutors = append(classTutors, "prof")
	fmt.Println(classTutors)

	getSecondToLast := classTutors[1:]
	fmt.Println(getSecondToLast)

}
