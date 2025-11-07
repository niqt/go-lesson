package oop

import (
	"fmt"
	"math"
)

// Interface definition
type Geometry interface {
	Area() float64
	Perimeter() float64
}

// Rect struct
type Rect struct {
	width, height float64
}

// Circle struct
type Circle struct {
	radius float64
}

// Rect implements Geometry interface
func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

// Circle implements Geometry interface
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Function that accepts any Geometry
func measure(g Geometry) {
	fmt.Printf("Area: %.2f\n", g.Area())
	fmt.Printf("Perimeter: %.2f\n", g.Perimeter())
}

func use_interface() {
	r := Rect{width: 3, height: 4}
	c := Circle{radius: 5}

	fmt.Println("Rectangle:")
	measure(r)

	fmt.Println("\nCircle:")
	measure(c)
}
