package main

import (
	"strings"
)

type Board [8][8]Piece

func (b *Board) At(sq Sq) Piece {
	return b[sq.Y][sq.X]
}

func (b *Board) AtCoord(s string) (Piece, Sq) {
	sq := Coord(s)
	return b.At(sq), sq
}

func (b *Board) Set(sq Sq, p Piece) {
	b[sq.Y][sq.X] = p
}

var InitialBoard string = strings.TrimSpace(`
rnbqkbnr
pppppppp
________
________
________
________
PPPPPPPP
RNBQKBNR
`)

func NewBoard() (b *Board) {
	b = &Board{}
	for y, row := range strings.Fields(InitialBoard) {
		for x, r := range row {
			b[y][x] = PieceFrom(r)
		}
	}
	return
}
