package main

import (
	"fmt"
	"testing"
)

func TestPiece_Slide(t *testing.T) {
	b := NewBoard()
	queen, _ := b.AtCoord("d1")
	moves := queen.Moves(Coord("e4"), b)
	fmt.Println(moves)
}
