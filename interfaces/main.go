package main

import "fmt"

type bot interface {
	getGreating() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreating(eb)
	printGreating(sb)
}

func (englishBot) getGreating() string {
	return "Hello"
}

func (spanishBot) getGreating() string {
	return "Hola"
}

func printGreating(b bot) {
	fmt.Println(b.getGreating())
}
