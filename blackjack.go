package main

func main() {
	deck := CreateDeck()
	player := Player{"Me", []Card{}}
	dealer := Player{"Dealer", []Card{}}
	deck.deal(&player, &dealer)
	player.displayHand()
	dealer.displayHand()
}
