package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	me := person{firstName: "Luca", lastName: "Lanziani"}
	fmt.Println(me)
	var alex person
	fmt.Println(alex)
}
