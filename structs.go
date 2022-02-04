package main

import (
	"fmt"
)

// BASICLY STRUCT AND STRUCT METHODS

type customer struct {
	name    string
	surname string
	age     int
}

func (c customer) save() {
	fmt.Println("Eklendi", c.name, c.surname, c.age)
}

// CREATE DEFAULT PARAMETER BY STRUCTS

type FromToNums struct {
	from, to int
}

func EvenNumbers(p FromToNums) {
	if p.from%2 != 0 {
		p.from++
	}

	if p.to == 0 {
		p.to = 10
	}

	for i := p.from; i < p.to; i += 2 {
		fmt.Println("Even Number : ", i)
	}
}

func main() {
	newCustomer := customer{"Uğur", "AKÇIL", 33}
	newCustomer.save()

	EvenNumbers(FromToNums{})
}
