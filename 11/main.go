// var x map[string]int is a map of strings to ints
/*x := make(map[string]int)
x["key"] = 10
fmt.Println(x["key"]*/
package main

import "fmt"

func main() {
	elements := make(map[string]string)

	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	fmt.Println(elements["O"])
}
