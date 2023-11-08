package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

// Create a new deck of cards and return it as a deck type (slice of strings)
func NewDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}

	return cards
}

// Print all the cards in the deck
func (cards Deck) Print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// Deal a hand of cards
func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

// Convert the deck to a string
func (d Deck) ToString() string {
	return strings.Join([]string(d), ", ")
}

// Save the deck to a file
func (d Deck) SaveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.ToString()), 0666)
}

// Read the deck from a file
func NewDeckFromFile(filename string) Deck {
	bs, error := os.ReadFile(filename)
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	} else {
		s := strings.Split(string(bs), ", ")
		return Deck(s)
	}

	return nil
}

// Shuffle the deck
func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
