package main

import "fmt"

func main() {
	x := 1.5
	y := square(&x)
	fmt.Println(y)
}

func square(x *float64) float64 {
	*x = *x * *x
	return *x
}
