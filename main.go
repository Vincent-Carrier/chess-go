package main

import "fmt"

func main() {
	s := &State{Board: NewBoard(), GameOver: false}
	p1, p2 := NewHuman(White), NewHuman(Black)
	active, passive := p1, p2
	for !s.GameOver {
	prompt:
		move := active.Prompt(active.Color, s)
		if !move.IsValid(s) {
			fmt.Println("Invalid move")
			goto prompt
		}
		fmt.Printf("%+v\n", move)
		s.Execute(move)
		active, passive = passive, active
	}
}
