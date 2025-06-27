package wasm

import (
	"fmt"
	"github.com/meteormin/slide-puzzle/internal/core"
)

const (
	DomIDPrefix = "tile-"
)

type Tile struct {
	DomID        string `json:"domId"`
	DisplayValue int    `json:"displayValue"`
	Empty        bool   `json:"fill"`
}

type EmptyXY struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Board struct {
	Size  int      `json:"size"`
	Tiles [][]Tile `json:"tiles"`
	Empty EmptyXY  `json:"empty"`
}

func (b *Board) MoveableDirection(x, y int) core.Direction {
	if x < 0 || y < 0 || x >= b.Size || y >= b.Size {
		return core.NONE
	}

	if x == b.Empty.X && y == b.Empty.Y {
		return core.NONE
	}

	if (x == b.Empty.X && (y == b.Empty.Y-1 || y == b.Empty.Y+1)) ||
		(y == b.Empty.Y && (x == b.Empty.X-1 || x == b.Empty.X+1)) {
		if y < b.Empty.Y {
			return core.Up
		} else if y > b.Empty.Y {
			return core.Down
		} else if x < b.Empty.X {
			return core.Left
		} else if x > b.Empty.X {
			return core.Right
		}
	}

	return core.NONE
}

func NewBoard(tiles [][]int) *Board {
	size := len(tiles)
	tileElements := make([][]Tile, size)
	empty := EmptyXY{}
	for i := 0; i < size; i++ {
		tileElements[i] = make([]Tile, size)
		for j := 0; j < size; j++ {
			element := Tile{}
			if tiles[i][j] == 0 {
				element.Empty = true
				empty.X = j
				empty.Y = i
			} else {
				element.DisplayValue = tiles[i][j]
				element.DomID = fmt.Sprintf("%s-%d-%d", DomIDPrefix, i, j)
			}
			tileElements[i][j] = element
		}
	}

	return &Board{
		Size:  size,
		Tiles: tileElements,
		Empty: empty,
	}
}
