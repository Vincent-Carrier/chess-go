package main

import "fmt"

func main() {
	b := NewBoard()
	p, sq := b.AtCoord("e2")
	fmt.Println(p.Moves(sq, b))
}
