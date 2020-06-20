package main

type Opponent struct {
	Color
	Player
}

type Player interface {
	Prompt(c Color, s *State) Move
}
