package main

import "fmt"

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}

func zero(xPtr *int) {
	*xPtr = 0
}

/*  * is also used to “dereference” pointer variables. Dereferencing a pointer gives us access to the value the
pointer points to. When we write *xPtr = 0 we are saying “store the int 0 in the memory location xPtr refers
to”. If we try xPtr = 0 instead we will get a compiler
error because xPtr is not an int it's a *int, which can
only be given another *int  */
