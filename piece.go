package main

type PieceKind int

//go:generate stringer -type=PieceKind
const (
	Empty PieceKind = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	Kind  PieceKind
	Color Color
}
