package main

import (
	"fmt"
	"time"
)

func saveLog() {
	fmt.Println("Log saved")
}

func main() {
	defer saveLog()
	m := 1
	for m <= 5 {
		fmt.Printf("Welcome %d times.\n", m)
		m = m + 1
	}

	for i := 6; i <= 10; i++ {
		fmt.Printf("Welcome 2.0 %d times\n", i)
	}

	d := time.Now()
	yil, ay, gun := d.Date()
	fmt.Println(yil, ay, gun)
}
