package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

type rectangle struct {
	length float64
	width  float64
}

func main() {
	s := square{side: 10}
	r := rectangle{length: 2, width: 4}
	t := triangle{base: 10, height: 5}

	printArea(s)
	printArea(r)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return ((0.5 * t.base) * t.height)
}

func (s square) getArea() float64 {
	return (s.side * s.side)
}

func (r rectangle) getArea() float64 {
	return (r.length * r.width)
}
