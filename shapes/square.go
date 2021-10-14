package main

type square struct {
	side float32
}

func (s square) area() float32 {
	return s.side * s.side
}
