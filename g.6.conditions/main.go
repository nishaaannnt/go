package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "hello, good morning"
	// string functions

	// if else
	if strings.Contains(greeting, "hello!") {
		fmt.Println("true bro")
	} else if 12 < 2 {
		fmt.Println("sad bro else if")
		// continue - in for loop
	} else {
		fmt.Println("sad bro else")
		// break - in for loop
	}

}
