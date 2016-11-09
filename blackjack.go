package main

import "fmt"

func main() {
	deck := CreateDeck()
	player := Player{"Me", []Card{}}
	dealer := Player{"Dealer", []Card{}}
	deck.deal(&player, &dealer)
	fmt.Printf("The player's hand is %v", player.Hand)
	fmt.Printf("The dealer's hand is %v", dealer.Hand)
}
