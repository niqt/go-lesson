package main

import "fmt"

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func measure(g Geometry) {
	fmt.Println("*************")
	fmt.Println("Geometry = ", g)
	fmt.Println("Area = ", g.Area())
	fmt.Println("Perimeter = ", g.Perimeter())
	fmt.Println("*************")
}

func InterfaceExample() {
	r := Rect{3, 4}
	c := Circle{5}
	measure(r)
	measure(c)
}
