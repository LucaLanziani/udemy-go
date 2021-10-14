package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	me := person{
		firstName: "Luca",
		lastName:  "Lanziani",
		contact: contactInfo{
			email:   "luca@lanziani.com",
			zipCode: 12345,
		},
	}
	fmt.Printf("%+v\n", me)

}
