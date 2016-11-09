package main

import "fmt"

type Player struct {
	Name string
	Hand []Card
}

func (p *Player) displayHand() {
	sub := cardSubString(p.Hand)
	fmt.Printf("%s has a hand with %s", p.Name, sub)
}

func cardSubString(h []Card) string {
	cards := ""
	d := ""
	for _, v := range h {
		d = v.display()
		cards = cards + d
	}
	return cards
}
