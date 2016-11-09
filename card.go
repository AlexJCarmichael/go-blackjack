package main

import (
	"fmt"
	"strconv"
)

type Card struct {
	Suit  string
	Value int
}

func (c *Card) display() string {
	v := valueDisplay(c.Value)
	return fmt.Sprintf("%s of %s ", v, c.Suit)
}

func valueDisplay(v int) string {
	if v == 11 {
		return "Jack"
	} else if v == 12 {
		return "Queen"
	} else if v == 13 {
		return "King"
	} else if v == 14 {
		return "Ace"
	} else {
		return strconv.Itoa(v)
	}
}
