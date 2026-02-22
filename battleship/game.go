package main

import "errors"

type Coord struct {
	X, Y int // x == column, y == row
}

type Hit bool

type Ship struct {
	Size   int
	Coords map[Coord]Hit
	Sunk   bool
}

const BoardSize = 10 // 10 x 10 row and width
type Board struct {
	Grid    []int // 0 == empty, 1 == ship
	Guesses []bool
	Ships   []*Ship
}

func GetBoard() *Board {
	grid := make([]int, BoardSize*BoardSize)
	guesses := make([]bool, BoardSize*BoardSize)
	return &Board{
		Grid:    grid,
		Guesses: guesses,
		Ships:   []*Ship{},
	}
}

func GetShip(size int) Ship {
	return Ship{
		Size:   size,
		Coords: make(map[Coord]Hit),
		Sunk:   false,
	}
}

func (b *Board) PlaceShip(ship Ship, startPos Coord, horizontal bool) error {
	shipCoords := GetShipCoords(startPos, ship.Size, horizontal)
	if b.IsValidSpace(shipCoords) {
		for _, c := range shipCoords {
			b.Grid[coordToIndex(c)] = 1
			ship.Coords[c] = false
		}
		b.Ships = append(b.Ships, &ship)
		return nil
	} else {
		return errors.New("Invalid place for ship")
	}
}

func coordToIndex(c Coord) int {
	return (c.Y * BoardSize) + c.X
}

func indexToCoord(index int) Coord {
	return Coord{
		Y: index / BoardSize,
		X: index % BoardSize,
	}
}

func GetShipCoords(startPos Coord, size int, horizontal bool) []Coord {
	coords := make([]Coord, size)
	for i := 0; i < size; i++ {
		if horizontal {
			coords[i] = Coord{X: startPos.X + i, Y: startPos.Y}
		} else {
			coords[i] = Coord{X: startPos.X, Y: startPos.Y + i}
		}
	}
	return coords
}

// check if coords are valid
// has to be within bounds of grid
// doesn't overlap with existing ships
func (b *Board) IsValidSpace(coords []Coord) bool {
	for _, c := range coords {
		if c.X >= BoardSize || c.Y >= BoardSize || c.X < 0 || c.Y < 0 {
			return false
		}
		if b.Grid[coordToIndex(c)] == 1 {
			return false
		}
	}
	return true
}

// mark coord as guessed
// check if ship exists at guess coord
func (b *Board) CheckIfMissileHit(c Coord) bool {
	idx := coordToIndex(c)
	if b.Guesses[idx] {
		return false
	}
	b.Guesses[idx] = true
	if b.Grid[idx] == 1 {
		return true
	}
	return false
}

func (b *Board) IsGameOver() bool {
	for _, s := range b.Ships {
		if !s.Sunk {
			return false
		}
	}
	return true
}
