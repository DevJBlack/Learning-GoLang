package main

import "fmt"

// var card string // Can create the variable out here but can't assign it to anything

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // Another way of writing the code above
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
