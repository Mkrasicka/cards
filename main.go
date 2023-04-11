package main

func main() {
	// cards := newDeck()

	// cards.print()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// greeting := "hello"
	// fmt.Println([]byte(greeting))

	// cards := newDeckFromFile("my_cards")
	// cards.print()
	cards := newDeck()
	cards.shuffleCards()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
