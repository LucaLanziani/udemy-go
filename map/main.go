package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
		"white": "#000000",
		"black": "#000000",
	}
	colors["white"] = "#FFFFFF"

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
