package main

func (p Piece) PawnMoves(from Sq, b *Board) (sqs Squares) {
	cardinal := p.Color.Forward()
	if fwd1, ok := p.SlideOne(from, cardinal, b); ok && b.At(fwd1).Color != p.Color.Opposite() {
		sqs = append(sqs, fwd1)
		if fwd2, ok := p.SlideOne(fwd1, cardinal, b); ok && b.At(fwd2).Color != p.Color.Opposite() {
			sqs = append(sqs, fwd2)
		}
	}
	for _, x := range [2]int{-1, 1} {
		cardinal.X = x
		if sq, ok := p.SlideOne(from, cardinal, b); ok && b.At(sq).Color == p.Color.Opposite() {
			sqs = append(sqs, sq)
		}
	}
	return
}
