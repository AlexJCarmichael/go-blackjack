package main

import "fmt"

func main() {
	deck := CreateDeck()
	fmt.Printf("The first card is %v", deck.Cards[0])
}
