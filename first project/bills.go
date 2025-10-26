package main

import "fmt"

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// function to make new bill 
func newBill(name string) bill {
  b := bill{
	name: name,
	items: map[string]float64{"cake": 5.5, "pie": 9.99},
	tip: 0,
  }

  return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	for k,v := range b.items {
		fs += fmt.Sprintf("%v ... $%v \n", k+":" , v)
		total += v
	}
    //total
	fs += fmt.Sprintf("%v ... %0.2f", "Total:", total)

	return fs
}