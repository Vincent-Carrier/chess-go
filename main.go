package main

import "fmt"

func main() {
	s := &State{
		Board: NewBoard(),
	}
	p1, p2 := &Opponent{White, Human{}}, &Opponent{Black, Human{}}
	active, passive := p1, p2
	gameOver := false
	for !gameOver {
	prompt:
		move := active.Prompt(active.Color, s)
		if !move.IsValid(s) {
			fmt.Println("Invalid move")
			goto prompt
		}
		active, passive = passive, active
	}
}
