package main

type State struct {
	Board *Board
	GameOver bool
}

func (s *State) Execute(m Move) {
	p := s.Board.Set(m.From, Piece{})
	s.Board.Set(m.To, p)
}
