package main

import (
	"math/rand"
	"time"
)

var Suites = [4]string{"Hearts", "Diamonds", "Clubs", "Spades"}

const MaxCardValue = 14

type Deck struct {
	Cards []Card
}

func CreateDeck() Deck {
	d := newDeck()
	d = shuffleDeck(d)
	deck := Deck{d}
	return deck
}

func newDeck() []Card {
	rand.Seed(time.Now().UTC().UnixNano())
	deck := []Card{}
	for _, v := range Suites {
		for c := 1; c <= MaxCardValue; c++ {
			card := Card{v, c}
			deck = append(deck, card)
		}
	}
	return deck
}

func shuffleDeck(d []Card) []Card {
	shuffled := 0
	for shuffled < 3 {
		for c := 0; c < MaxCardValue; c++ {
			randCard := rand.Intn(len(d))
			d[c], d[randCard] = d[randCard], d[c]
		}
		shuffled++
	}
	return d
}
