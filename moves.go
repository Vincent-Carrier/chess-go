package main

func makeRotations(v Vec) (rotations [4]Vec) {
	for i := 0; i < 4; i++ {
		v = v.Rotate()
		rotations[i] = v
	}
	return
}

func combineCardinals(cardinals ...Vec) (rotations []Vec) {
	for _, cardinal := range cardinals {
		rot := makeRotations(cardinal)
		rotations = append(rotations, rot[:]...)
	}
	return
}

func (p Piece) SlideOne(from Sq, towards Vec, b *Board) (Sq, bool) {
	target := from.Add(towards)
	if !target.IsWithinBoard() || b.At(target).Color == p.Color {
		return target, false
	}
	return target, true
}

func (p Piece) Slide(from Sq, towards Vec, b *Board) (sqs Squares) {
	for {
		if sq, ok := p.SlideOne(from, towards, b); ok {
			sqs = append(sqs, sq)
			if p.Color == b.At(sq).Color.Opposite() {
				return
			}
			from = sq
		} else {
			return
		}
	}
}

func (p Piece) shortMoves(from Sq, b *Board, cardinals ...Vec) (moves Squares) {
	for _, cardinal := range cardinals {
		if sq, ok := p.SlideOne(from, cardinal, b); ok {
			moves = append(moves, sq)
		}
	}
	return
}

func (p Piece) longMoves(from Sq, b *Board, cardinals ...Vec) (moves Squares) {
	for _, towards := range cardinals {
		line := p.Slide(from, towards, b)
		moves = append(moves, line...)
	}
	return
}

var (
	knightCardinals []Vec = combineCardinals(Vec{1, 2}, Vec{2, 1})
	bishopCardinals []Vec = combineCardinals(Vec{1, 1})
	rookCardinals   []Vec = combineCardinals(Vec{1, 0})
	queenCardinals  []Vec = combineCardinals(Vec{1, 1}, Vec{1, 0})
)

func (p *Piece) Moves(from Sq, b *Board) (sqs Squares) {
	switch p.Kind {
	case Pawn:
		return p.PawnMoves(from, b)
	case Knight:
		return p.shortMoves(from, b, knightCardinals...)
	case Bishop:
		return p.longMoves(from, b, bishopCardinals...)
	case Rook:
		return p.longMoves(from, b, rookCardinals...)
	case Queen:
		return p.longMoves(from, b, queenCardinals...)
	case King:
		return p.shortMoves(from, b, queenCardinals...)
	default:
		panic("Nope...")
	}
}
