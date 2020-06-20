package main

type Color int

const (
	None Color = iota
	Black
	White
)

func (c Color) Opposite() Color {
	switch c {
	case White:
		return Black
	case Black:
		return White
	default:
		return None
	}
}

func (c Color) Forward() Vec {
	switch c {
	case White:
		return Vec{0, -1}
	case Black:
		return Vec{0, 1}
	default:
		panic("Called Forward on None Color")
	}
}
