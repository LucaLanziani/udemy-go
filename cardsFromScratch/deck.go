package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}
	cardSuits := []string{"Spades", "Diamons", "Hearts", "Clubs"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	by, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(by), ","))
}

func (d deck) shuffle() {
	rand.Seed(time.Now().Unix())

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
