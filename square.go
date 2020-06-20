package main

type Sq Vec
type Squares []Sq

func Coord(s string) Sq {
	file, rank := int(s[0]), int(s[1])
	return Sq{
		file - 97,
		56 - rank,
	}
}

func (sq Sq) IsWithinBoard() bool {
	return sq.X >= 0 && sq.X < 8 && sq.Y >= 0 && sq.Y < 8
}

func (sq Sq) Add(v Vec) Sq {
	return Sq{
		sq.X + v.X,
		sq.Y + v.Y,
	}
}
