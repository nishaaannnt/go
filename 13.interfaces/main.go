package main

import (
	"fmt"
	"math"
	"reflect"
)

// Interface
// can be used to create general functions

// way to group together types
type shape interface {
	// based on functions associated to those types
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func printShapeDetails(s shape) {
	fmt.Println("Area of", reflect.TypeOf(s), "is", s.area())
	fmt.Println("Circumference of", reflect.TypeOf(s), "is", s.circumf())
}

func main() {
	// this is a slice of interface shape
	shapes := []shape{
		square{length: 20},
		circle{radius: 10},
		square{length: 40},
	}

	for _, v := range shapes {
		printShapeDetails(v)
	}
}
