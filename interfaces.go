package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c rectangle) area() float64 {
	fmt.Println("Calculating rectangle area..")
	return c.width * c.height
}

func (c circle) area() float64 {
	fmt.Println("Calculating circle area..")
	return math.Pi * c.radius * c.radius
}

func areaPrint(s shape) {
	fmt.Println("By the interface : ", s.area())
}

func main() {
	rec := rectangle{width: 10, height: 5}
	cir := circle{radius: 5}

	fmt.Println(rec.area())
	fmt.Println(cir.area())

	areaPrint(rec)
	areaPrint(cir)

	shapes := []shape{
		rectangle{width: 5, height: 4},
		circle{radius: 6},
		rectangle{width: 8, height: 5},
		circle{radius: 8},
	}

	for _, s := range shapes {
		areaPrint(s)
	}
}
