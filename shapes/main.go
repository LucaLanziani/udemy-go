package main

import "fmt"

type square struct {
	side float32
}

type triangle struct {
	base   float32
	height float32
}

func (t triangle) getArea() float64 {
	return float64(t.base * t.height / 2)
}

func (s square) getArea() float64 {
	return float64(s.side * s.side)
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	printArea(square{4})
	printArea(triangle{2, 2})
}
