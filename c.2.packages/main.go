package main

import "fmt"

func main() {
	// fmt - format
	age := 35
	// print
	fmt.Print("hello")
	fmt.Print("hello\n")
	// adds '\n' automatically
	fmt.Println("hello, you age is", age)

	// Printf (formatted strings) %_ = format specifier
	// %v - variable
	fmt.Printf("my age is %v \n", age)
	// provides type
	fmt.Printf("my age is of type %T \n", age)

	//
	fmt.Printf("my age is  %2f \n", 25.5555)
	fmt.Printf("my age is  %f \n", 25.5555)

	// Sprintf - save formatted string
	str := fmt.Sprintf("my age is %v", age)

	fmt.Println("saved string is", str)
}
