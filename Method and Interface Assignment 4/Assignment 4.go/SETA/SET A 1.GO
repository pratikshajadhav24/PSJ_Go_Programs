//1. Write a program in go language to create an interface shape that includes area and perimeter. Implements these methods in circle and rectangle type.

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length, Width float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func main() {
	var s Shape

	s = Circle{Radius: 5}
	fmt.Println("Circle Area:", s.Area())
	fmt.Println("Circle Perimeter:", s.Perimeter())

	s = Rectangle{Length: 4, Width: 6}
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())
}
