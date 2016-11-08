package main

var Suites = [4]string{"Hearts", "Diamonds", "Clubs", "Spades"}

const MaxCardValue = 14

type Deck struct {
	Cards []Card
}

func CreateDeck() Deck {
	d := newDeck()
	deck := Deck{d}
	return deck
}

func newDeck() []Card {
	deck := []Card{}
	for _, v := range Suites {
		for c := 1; c <= MaxCardValue; c++ {
			card := Card{v, c}
			deck = append(deck, card)
		}
	}
	return deck
}
