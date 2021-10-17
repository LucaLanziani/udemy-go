package main

import "fmt"

func main() {
	const day int = 7
	fmt.Println("test", day)

	const secondsInHours = 3600

	durantion := 234

	fmt.Printf("Durantion in seconds %v\n", durantion*secondsInHours)

	const (
		one = 3
		two
		three = 4
	)

	fmt.Println(one, two, three)
	types()
	iotaExample()
}
