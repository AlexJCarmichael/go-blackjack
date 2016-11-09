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
		for c := 2; c <= MaxCardValue; c++ {
			card := Card{v, c}
			deck = append(deck, card)
		}
	}
	return deck
}

func shuffleDeck(d []Card) []Card {
	for shuffled := 0; shuffled < 10; shuffled++ {
		for c := 0; c < MaxCardValue; c++ {
			randCard := rand.Intn(len(d))
			d[c], d[randCard] = d[randCard], d[c]
		}
	}
	return d
}

func (d *Deck) draw(p *Player) {
	var card Card
	card, d.Cards = d.Cards[len(d.Cards)-1], d.Cards[:len(d.Cards)-1]
	p.Hand = append(p.Hand, card)
}

func (d *Deck) deal(p *Player, dealer *Player) {
	for i := 0; i < 2; i++ {
		d.draw(p)
		d.draw(dealer)
	}
}
