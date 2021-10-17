package main

import (
	"fmt"
)

func iotaExample() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3)

	const (
		c11 = iota
		c22
		c33
	)

	fmt.Println(c11, c22, c33)

	const (
		a = (iota * 2) + 1
		b
		c
	)

	fmt.Println(a, b, c)

	// x = -2, y = -4, z = -5
	const (
		x = -(iota + 2)
		_
		y
		z
	)

	fmt.Println(x, y, z)
}
