package main

import (
	"fmt"
	"time"
)

func EvenNumbersSum(total chan int) {
	time.Sleep(time.Second * 3)
	fmt.Println("Even Numbers init..")
	var to int = 10
	var sum int

	for i := 0; i <= to; i += 2 {
		sum += i
	}

	total <- sum
}

func OddNumbersSum(total chan int) {
	time.Sleep(time.Second * 3)
	fmt.Println("Odd Numbers init..")
	var to int = 10
	var sum int

	for i := 1; i <= to; i += 2 {
		sum += i
	}

	total <- sum
}

func main() {
	fmt.Println("All process will be init..")
	evenSumCn, oddSumCn := make(chan int), make(chan int)
	fmt.Println("Channels created..")

	go EvenNumbersSum(evenSumCn)
	go OddNumbersSum(oddSumCn)

	fmt.Println("Routines created..")

	evenSum, oddSum := <-evenSumCn, <-oddSumCn

	fmt.Println("Getting channel vars..")

	multiplication := evenSum * oddSum

	fmt.Println("Even Numbers Sum, Odd Numbers Sum : ", evenSum, ",", oddSum)
	fmt.Println("Multiplication : ", multiplication)

	fmt.Println("The End..")
}
