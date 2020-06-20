package main

type MoveKind int

const (
	Normal = iota
)

type Move struct {
	Kind MoveKind
	From Sq
	To Sq
	Capture
}

type Capture struct {
	Piece
	Sq
}

func (m Move) IsValid(s *State) bool {
	return true
}
