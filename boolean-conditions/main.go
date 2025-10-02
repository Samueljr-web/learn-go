package main

import "fmt"


func main() {
	age := 45

	fmt.Println(age >= 50)
	fmt.Println(age <= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	}else if age < 45 {
		fmt.Println("age is less than 45")
	}else {
		fmt.Println("age is not less than 45")
	}

	names := []string{"sam", "ad", "odun", "ox"}

	for index, value := range names{
		if index == 1 {
			fmt.Println("continuing at index", index)
			continue
		}
		if (index > 2){
			fmt.Printf("breaking at pos %v", index)
			break
		}
		fmt.Printf("the value for postioning %v and %v \n", value, index)
	}

}
