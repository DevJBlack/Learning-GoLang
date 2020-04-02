package main

// var card string // Can create the variable out here but can't assign it to anything

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // Another way of writing the code above
	// card := newCard()
	// fmt.Println(card)
	// cards := newDeck()

	// hand, remaingCards := deal(cards, 5)

	// hand.print()
	// remaingCards.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
