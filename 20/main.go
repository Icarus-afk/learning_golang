package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	z float64
}

func main() {
	c := Circle{x: 0, y: 0, z: 5}
	fmt.Println(c.x, c.y, c.z)
	fmt.Println(c.area())
}

func (c *Circle) area() float64 {
	return math.Pi * c.z * c.z
}
