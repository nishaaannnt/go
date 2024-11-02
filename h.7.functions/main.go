package main

import (
	"fmt"
	"math"
	"strings"
)

// simple function
func sayGreeting(s string) {
	fmt.Println("Hello", s)
}

// pass function in a function
func loopName(n []string, f func(string)) {
	for _, val := range n {
		f(val)
	}
}

// returning a value from function
func circle(r float64) float64 {
	return math.Pi * r * r
}

// return multiple value
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, val := range names {
		initials = append(initials, val[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	greeting := "hello, good morning"
	// string functions
	sayGreeting(greeting)

	loopName([]string{"nishant", "dixit", "sarvesh", "hare krishna"}, sayGreeting)

	fmt.Println(circle(4))
	fmt.Printf("circle is %0.3f \n", circle(4))

	firstName, lastName := getInitials("nishant dixit")
	fmt.Println(firstName, lastName)
}
