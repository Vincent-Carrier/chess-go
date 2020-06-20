package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoord(t *testing.T) {
	a1, e1 := Coord("e4"), Sq{4, 4}
	a2, e2 := Coord("h1"), Sq{7, 7}
	a3, e3 := Coord("a8"), Sq{0, 0}
	assert.Equal(t, e2, a2)
	assert.Equal(t, e1, a1)
	assert.Equal(t, e3, a3)
	fmt.Println([]Sq{e1, e2})
}
