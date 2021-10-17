package main

import "fmt"

func types() {
	const x = 2
	const y = 2.2 * 5

	fmt.Printf("%T\n", y)
	var i int = x
	var j float64 = x
	var p byte = x

	fmt.Println(i, j, p)

}
