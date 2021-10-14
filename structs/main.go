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
	alex.firstName = "Alex"
	alex.lastName = "None"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
