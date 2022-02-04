package main

import (
	"fmt"
	"sort"
)

func main() {

	harfler := []string{"U", "Ğ", "U", "R", "A", "K", "Ç", "I", "L"}

	sort.Strings(harfler)
	fmt.Println(harfler)

	sayilar := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}

	sortSearch := sort.Search(len(sayilar), func(i int) bool { return sayilar[i] >= 2 })

	fmt.Println(sortSearch)
	fmt.Println(sayilar[sortSearch])

	if sortSearch < len(sayilar) && sayilar[sortSearch] == 2 {
		fmt.Println("there is")
	} else {
		fmt.Println("there isnt")
	}
}
