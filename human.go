package main

import (
	"fmt"
	"io"
	"os"
)

type Human struct {
	r io.Reader
	w io.Writer
}

func NewHuman(c Color) *Opponent {
	p := &Human{
		r: os.Stdin,
		w: os.Stdout,
	}

	return &Opponent{Color: c, Player: p}
}

func (h Human) Prompt(c Color, s *State) Move {
	fmt.Fprintf(h.w, "%s's player turn to play", c)
	var str string
	fmt.Fscanf(h.r, "%c%d\n", &str)
	return ParseMove(&str)
}

func ParseMove(s *string) (m Move) {

	return
}
