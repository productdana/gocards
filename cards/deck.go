package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new deck type
// which is a slice of type string
// go back and replace all the []string with deck
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Dimaonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// []string(d) is slice of strings
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// bs is for byte slice
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option 1: log error and return a call to newDeck()
		// option 2: log error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// s is a slice of strings
	s := strings.Split(string(bs), ",") // string(bs) converts byte slice to string: "Ace of Spades,Two of Spades,Three of Spades,..."
	// convert slice of strings to deck type so it has access to other methods
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source) // r is Type of Rand; random number generator

	// i is index; not necessary to include element
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// swap cards - d[i] is assigned to the first thing after = and vice versa
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
