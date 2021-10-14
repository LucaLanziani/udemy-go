package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	me := person{
		firstName: "Luca",
		lastName:  "Lanziani",
		contactInfo: contactInfo{
			email:   "luca@lanziani.com",
			zipCode: 12345,
		},
	}
	me.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
