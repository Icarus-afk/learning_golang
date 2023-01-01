package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	height, width float64
}
type Circle struct {
	radius float64
}
type Shape interface {
	area() float64
}

func main() {
	var r, h, w float64
	fmt.Println("Enter circle radius: ")
	fmt.Scan(&r)
	fmt.Println("Enter rectange height and width: ")
	fmt.Scan(&h, &w)

	x1 := Rectangle{h, w}
	x2 := Circle{r}

	fmt.Println("Rectangle height: ", x1.height)
	fmt.Println("Rectangle Width: ", x1.width)
	fmt.Println("Rectangle area: ", x1.area())
	fmt.Println("Circle Radius: ", x2.radius)
	fmt.Println("Circle area: ", x2.area())
	fmt.Println("Total area: ", totalArea(&x1, &x2))

}

func (r *Rectangle) area() float64 {
	x := r.height * r.width
	return x
}

func (c *Circle) area() float64 {
	x := c.radius * c.radius * math.Pi
	return x
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
