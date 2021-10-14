package main

import "math"

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return c.radius * c.radius * math.Pi
}
