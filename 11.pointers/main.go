package main

import (
	"fmt"
)

// *type -> pointer to a type of val
// * is dereference operator
func updateName(x *string) {
	// get the value in the pointer and updates it
	*x = "dixit"
}

func main() {
	name := "nishant"
	// fmt.Println("name's memory address", &name) // name's memory address 0x14000010040
	m := &name
	// *m -> provides value at that mem address
	fmt.Println("m is", m, "and its value is", *m) // m is 0x14000010050 and its value is nishant

	fmt.Println("before call, name was ", name)
	updateName(&name)
	fmt.Println("name now is", name)
}
