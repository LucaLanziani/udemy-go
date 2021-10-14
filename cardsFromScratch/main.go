package main

func main() {
	deck := newDeckFromFile("my_deck")
	deck.shuffle()
	deck.print()
	// deck.saveToFile("my_deck")
}
