package main

import (
	"strings"
	"unicode"
)

func (p Piece) String() string {
	var letter rune
	switch p.Kind {
	case Pawn:
		letter = '♙'
	case Knight:
		letter = '♘'
	case Bishop:
		letter = '♗'
	case Rook:
		letter = '♖'
	case Queen:
		letter = '♕'
	case King:
		letter = '♔'
	default:
		return " "
	}
	if p.Color == Black {
		letter += 6
	}
	return string(letter)
}

func PieceFrom(r rune) (p Piece) {
	if unicode.IsLower(r) {
		p.Color = Black
	} else if unicode.IsUpper(r) {
		p.Color = White
	} else {
		p.Color = None
	}

	switch unicode.ToLower(r) {
	case 'p':
		p.Kind = Pawn
	case 'n':
		p.Kind = Knight
	case 'b':
		p.Kind = Bishop
	case 'r':
		p.Kind = Rook
	case 'q':
		p.Kind = Queen
	case 'k':
		p.Kind = King
	default:
		p.Kind = Empty
	}
	return
}

func (b Board) String() string {
	var builder strings.Builder
	for _, row := range b {
		for _, p := range row {
			builder.WriteString(p.String())
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}

func (sqs Squares) String() string {
	var builder strings.Builder
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if sqs.Contains(Sq{j, i}) {
				builder.WriteRune('X')
			} else {
				builder.WriteRune('_')
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func (sqs Squares) Contains(square Sq) bool {
	for _, sq := range sqs {
		if sq == square {
			return true
		}
	}
	return false
}
