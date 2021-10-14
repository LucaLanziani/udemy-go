package main

import "fmt"

type shape interface {
	area() float32
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func main() {
	printArea(square{4})
	printArea(circle{2})
	printArea(triangle{2, 2})

}
