package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello, good morning"
	// string functions
	fmt.Println(strings.Contains(greeting, "hello"))

	// manipulating returns new strings
	fmt.Println(strings.ReplaceAll(greeting, "hello", "dixit"))
	// still show "hello, good morning"
	fmt.Println(greeting)

	arr := []int{1, 5, 2, 3, 9, 7, 8}

	// SORTING in GO
	sort.Ints(arr)
	fmt.Println(sort.SearchInts(arr, 5)) // can also search using sort
	fmt.Println(arr)

	// if else
	if strings.Contains(greeting, "hello!") {
		fmt.Println("true bro")
	} else {
		fmt.Println("sad bro")
	}

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println("yo", i)
	}
}
