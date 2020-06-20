package main

//go:generate stringer -type=MoveKind
type MoveKind int

const (
	Normal = iota
)

type Move struct {
	Kind MoveKind
	From Sq
	To Sq
	Capture *Capture
}

type Capture struct {
	Piece
	Sq
}

func (m Move) IsValid(s *State) bool {
	return true
}
