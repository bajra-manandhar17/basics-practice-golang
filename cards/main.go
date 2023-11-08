package main

import (
	"cards/deck"
)

func main() {
	cards := deck.NewDeckFromFile("my_cards")
	cards.Print()
}
