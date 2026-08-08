package main

import "fmt"

type geometry interface {
	area() float64
	parameter() float64
	getName() string
}

type circle struct {
	radius float64
	name   string
}

type square struct {
	side float64
	name string
}

type triangle struct {
	base   float64
	height float64
	name   string
}

func (c circle)	area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) parameter() float64 {
	return 2 * 3.14 * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) parameter() float64 {
	return 4 * s.side
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) parameter() float64 {
	// Simplified for this example - in reality, you'd need to calculate the actual perimeter
	return t.base + 2*t.height
}

func (c circle) getName() string {
	return "Name of geometry type: " + c.name
}

func (s square) getName() string {
	return "Name of geometry type: " + s.name
}

func (t triangle) getName() string {
	return "Name of geometry type: " + t.name
}

func main() {
	fmt.Println("Experiencing interfaces in Go!")

	t := triangle{base: 3, height: 4, name: "MyTriangle"}
	measure(t)

	c := circle{radius: 5, name: "MyCircle"}
	measure(c)

	s := square{side: 6, name: "MySquare"}
	measure(s)

	myPrinter(t.base, t.height, t.area(), t.parameter())
}

func measure (g geometry) {
	fmt.Println(g.getName())
	fmt.Println("Area:", g.area())
	fmt.Println("Parameter:", g.parameter())
}

func myPrinter(i ... interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}