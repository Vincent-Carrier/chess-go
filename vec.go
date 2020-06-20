package main

type Vec struct {
	X, Y int
}

func (v Vec) Times(n int) Vec {
	return Vec{v.X * n, v.Y * n}
}

func (v Vec) Rotate() Vec {
	return Vec{v.Y * -1, v.X}
}
