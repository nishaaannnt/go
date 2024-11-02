package main

import (
	"fmt"
)

func main() {
	// greeting := "hello, good morning"
	// string functions

	// // if else
	// if strings.Contains(greeting, "hello!") {
	// 	fmt.Println("true bro")
	// } else {
	// 	fmt.Println("sad bro")
	// }
	x := 0
	// for loop
	// similar to while
	for x < 5 {
		fmt.Println("x is - ", x)
		x++
	}

	// loop through an array
	str := []string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
	// another way
	for index, val := range str {
		fmt.Printf("at %v index has %v \n", index, val)
	}
}
