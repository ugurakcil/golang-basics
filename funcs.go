package main

import "fmt"

func main() {

	// Defer flag***********

	defer fmt.Println("tamam tamam")

	//**************************
	fmt.Println(bunlariTopla(3, 4, 5))

	// Closure funcs***********
	toplam := func(x, y int) int {
		return x + y
	}

	fmt.Println(toplam(3, 5))

	//**************************

	fmt.Println(fakto(3))

	//**************************
}

// Variadic funcs
func bunlariTopla(sayilar ...int) int {
	toplam := 0

	for _, sayi := range sayilar {
		toplam += sayi
	}

	return toplam
}

// Recursion funcs
func fakto(simdikiRakam uint) uint {
	if simdikiRakam == 0 {
		return 1
	}

	return simdikiRakam * fakto(simdikiRakam-1)
}
