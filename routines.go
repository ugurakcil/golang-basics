package main

import (
	"fmt"
	"time"
)

func EvenNumbers() {
	fmt.Print("\nEven Numbers : ")
	var to int = 10

	for i := 0; i <= to; i += 2 {
		fmt.Print(i)

		if i < to {
			fmt.Print(",")
		}
	}

	fmt.Print("\n")
}

func OddNumbers() {
	fmt.Print("\nOdd Numbers : ")
	var to int = 10

	for i := 1; i <= to; i += 2 {
		fmt.Print(i)

		if i < to-1 {
			fmt.Print(",")
		}
	}

	fmt.Print("\n")
}

func main() {
	go EvenNumbers()
	go OddNumbers()
	go EvenNumbers()
	go OddNumbers()
	time.Sleep(time.Second * 5)
}
