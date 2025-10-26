package main

import "fmt"

func main() {
	mybill := newBill("Samuel's bill")

    mybill.format()
	fmt.Println(mybill.format())
}