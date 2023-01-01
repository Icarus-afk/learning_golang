// factorial recursion
package main

import "fmt"

func main() {
	var x uint
	fmt.Println("Enter a number: ")
	fmt.Scanln(&x)

	fmt.Println(factorial(x))
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
