package main

type Kind int

const (
	Empty Kind = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	Kind  Kind
	Color Color
}
