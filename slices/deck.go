package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// creates equivalence between deck alias and []string type
type deck []string

func newDeck() deck {
	d := deck{}
	cardSuits := []string{
		"Spades",
		"Clubs",
		"Hearts",
		"Diamonds",
	}
	cardValues := []string {
		"Ace",
		"Two",
		"Three",
		"Four",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, value + " of " + suit)
		}
	}
	return d

}

// receiver function - any variable of the receiver type will have access to this method
// receiver arg is typically referred to by a single letter by convention
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i , card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return  d[0:handSize], d[handSize:]
}


func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFile(filename string) deck {
	byteDeck, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
	return deck(strings.Split(string(byteDeck), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}