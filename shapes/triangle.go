package main

type triangle struct {
	base   float32
	height float32
}

func (t triangle) area() float32 {
	return t.base * t.height / 2
}
