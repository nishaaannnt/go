package main

import (
	"fmt"
)

// MAPS in Golang
// maps allow to store key - value pairs

func main() {

	// map[typeof key] typeof value
	menu := map[string]float64{
		"soup":   400.9,
		"pie":    799,
		"toffee": 800,
	}

	fmt.Println("yo", menu)
	fmt.Println(menu["soup"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// int as key
	phoneBook := map[int]string{
		9372524596: "nishant",
		9867564586: "dixit", // requires end ,
	}
	phoneBook[9372524596] = "rishi" // updating
	fmt.Println(phoneBook[9372524596])
}
