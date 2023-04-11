package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create the new type of deck which is a slice of string
type deck []string

func newDeck() deck {
	// create list of cards
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		// var card string = cardSuit

		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
	// return list of cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// []string(d) converting byte/array into a string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToDeck(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// byte slice argument
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// return a message with an error
		fmt.Println("Found error:", err)
		// exit program
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	// deck is a type and we are passing s variable which is a slice of strings and deck type is a slice as well
	return deck(s)
}

// for each index , card in cards
// generate a random number between 0- len(cards)
// swap the current card and the cards at cards[randomNumber]

func (d deck) shuffleCards() {
	// to generate every time we start the program different int64 number
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		randomNumber := r.Intn(len(d) - 1)
		// swapping logic
		d[i], d[randomNumber] = d[randomNumber], d[i]
	}
}
