package main

import (
	"fmt"
)

// Go is a pass-by-value language
// when we pass val to function -> go creates copy of them

// copies the value of x and stores in another mem address
func updateName(x string) {
	x = "dixit" // cannot work
}

// they have pass by reference as pointer is copied
func updateMapValue(m map[string]float64) {
	m["soup"] = 200 // will work
}

func main() {
	// Non Pointer (group A) types -> strings, int, bool, float, array, structs
	// Group A have pass-by-value by default
	name := "nishant"
	updateName(name) // not work

	// --------------------------------------

	// Pointer Wrapped (group B) types -> slices, maps, functions
	// Group B have pass-by-reference

	menu := map[string]float64{
		"soup":   400.9,
		"pie":    799,
		"toffee": 800,
	}

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}
	// updates menu["soup"] = 200
	updateMapValue(menu)

	fmt.Println("\n\nupdated")

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

}
