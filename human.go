package main

import (
	"fmt"
)

type Human struct{}

func NewHuman(c Color) *Opponent {
	return &Opponent{
		Color:  c,
		Player: Human{},
	}
}

func (h Human) Prompt(c Color, s *State) Move {
	fmt.Println(s.Board)
	fmt.Printf("%s's player turn to play\n", c)
	var from, to string
	fmt.Scan(&from, &to)
	return ParseMove(from, to)
}

func ParseMove(from, to string) Move {
	return Move{
		Kind: Normal,
		From: Coord(from),
		To:   Coord(to),
	}
}
