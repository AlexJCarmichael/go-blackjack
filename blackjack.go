package main

import "fmt"

func main() {
	deck := CreateDeck()
	fmt.Printf("The first card is %v", deck.Cards[0])
	fmt.Printf("The second card is %v", deck.Cards[1])
	fmt.Printf("The third card is %v", deck.Cards[2])
	fmt.Printf("The last card is %v", deck.Cards[51])
}
