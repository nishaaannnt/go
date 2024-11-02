package main

// formatting strings and printing messages to console
import "fmt"

var someName = "yo"

// someName := "yo" - error

// entry point of an application
func main() {
	fmt.Println("Hello, World!")

	// I. STRINGSS
	// declaring var ways
	var nameOne string = "ram"
	var nameTwo = "krishna"

	// this has default ""
	var nameThree string // setting var for future

	fmt.Println(nameOne, nameTwo, nameThree)

	nameTwo = "hari"
	nameThree = "dev"
	//  nameTwo = 2  -> throws error

	fmt.Println(nameOne, nameTwo, nameThree)

	// common used way
	// NOTE - can only use this way in functions
	// outside the function - error
	nameFour := "prabhu"
	fmt.Println(nameFour)

	// II. ints
	var ageOne int = 12
	// - for knowledge
	// var age8 int8 = 1   // -128 to 127
	// var age16 int16 = 1 // -32k to 32k
	// var age32 int32 = 1
	// var age64 int64 = 1
	var age64 uint = 1 //only positive
	var ageTwo = 32

	// mostly used
	ageThree := 32

	fmt.Println(ageOne, ageTwo, ageThree, age64)

	// III. float
	var floOne float32 = 25.98

	// mostly used
	var floTwo float64 = 251123123123.98
	scoreThree := 1.5 // infered as float 64

	fmt.Println(floTwo, floOne, scoreThree)
}
