package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	// BASIC ARRAY AND FOR

	yillar := []int{1989, 1991, 2016}

	for yil := 1980; yil <= time.Now().Year(); yil++ {
		fmt.Println(yil)
	}

	// RANGE

	for _, yil := range yillar {
		fmt.Println("Important year : ", yil)
	}

	// MAP
	dogumHaritasi := map[string]int{"gokce": 911217, "ugur": 890808, "defne": 160905}

	for isim, dogumTarihi := range dogumHaritasi {
		fmt.Printf("%s -> %s \n", isim, dogumTarihi)
		fmt.Println(dogumTarihi)
	}

	// SLICE

	isimler := make([]string, 2)

	isimler[0] = "Uğur"
	isimler[1] = "Ayhan"

	fmt.Println(isimler)

	isimler = append(isimler, "Türkan")

	fmt.Println("after append : ", isimler)

	tumIsimler := make([]string, len(isimler))

	copy(tumIsimler, isimler)

	fmt.Println("after copy : ", tumIsimler)

	tumIsimler = append(tumIsimler, "Bahar", "Defne", "Hasan", "Fatma")
	fmt.Println("before DeleteOnSlice : ", tumIsimler)
	DeleteOnSlice(&tumIsimler, 3)
	fmt.Println(tumIsimler)
	DeleteOnSlice(&tumIsimler, 4)
	fmt.Println(tumIsimler)

	interfaceSlice := make([]interface{}, 2)

	interfaceSlice[0] = 4
	interfaceSlice[1] = "test2"

	fmt.Println(interfaceSlice)
}

func DeleteOnSlice(slice interface{}, item int) {
	//return append(slice[:delete], slice[delete+1:]...)
	vField := reflect.ValueOf(slice)
	value := vField.Elem()
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		result := reflect.AppendSlice(value.Slice(0, item), value.Slice(item+1, value.Len()))
		value.Set(result)
	}
}
