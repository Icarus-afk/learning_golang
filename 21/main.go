package main

import (
	"fmt"
)

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func main() {
	var a1, a2, a3, a4 float64
	fmt.Println("Enter x1 y1 x2 y2: ")
	fmt.Scanln(&a1)
	fmt.Scanln(&a2)
	fmt.Scanln(&a3)
	fmt.Scanln(&a4)

	r := Rectangle{a1, a2, a3, a4}

	fmt.Println("x1: ", r.x1)
	fmt.Println("y1: ", r.y1)
	fmt.Println("x2: ", r.x2)
	fmt.Println("y2: ", r.y2)
}
