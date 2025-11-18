package snake

import (
	"math/rand/v2"
)

type Coords struct{ X, Y int }

type Direction = int32

const (
	U Direction = iota
	D
	L
	R
)

var directionCoords = [4]Coords{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

type snake struct {
	m            map[Coords]bool
	Food         Coords
	Snake        []Coords // last element is snake mouth
	maxX, maxY   int
	Dir          Direction
	CanChangeDir bool
}

func NewSnake(mx, my int) *snake {
	snak := make([]Coords, 0, mx*my)
	snak = append(snak, Coords{mx / 2, my / 2})
	m := make(map[Coords]bool)
	m[Coords{mx / 2, my / 2}] = true
	d := U
	s := snake{Snake: snak, Food: Coords{rand.IntN(mx), rand.IntN(my)}, maxX: mx, maxY: my, Dir: d, m: m}
	return &s
}

func (s *snake) Next() bool {
	length := len(s.Snake)
	dir := s.Dir
	changed := directionCoords[dir]
	next := Coords{(s.Snake[length-1].X + changed.X) % s.maxX, (s.Snake[length-1].Y + changed.Y) % s.maxY}
	if next.X < 0 {
		next.X += s.maxX
	}
	if next.Y < 0 {
		next.Y += s.maxY
	}
	s.Snake = append(s.Snake, next)
	s.m[next] = true
	if s.Food == next {
		s.Food = Coords{rand.IntN(s.maxX), rand.IntN(s.maxY)}
		return len(s.Snake) == len(s.m)
	}
	delete(s.m, s.Snake[0])
	s.Snake = s.Snake[1:]
	s.CanChangeDir = true
	return len(s.Snake) == len(s.m)
}
