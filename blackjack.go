package main

import "fmt"

func main() {
	deck := CreateDeck()
	card := deck.draw()
	fmt.Printf("The drawn card is %v", card)
	fmt.Printf("The length of the deck is %v", len(deck.Cards))
}
