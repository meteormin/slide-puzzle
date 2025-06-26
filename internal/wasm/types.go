package wasm

import "strconv"

const (
	DomIDPrefix = "tile-"
)

type Tile struct {
	DomID        string `json:"domId"`
	DisplayValue int    `json:"displayValue"`
	Empty        bool   `json:"fill"`
}

type Board struct {
	Size  int      `json:"size"`
	Tiles [][]Tile `json:"tiles"`
}

func NewBoard(tiles [][]int) *Board {
	size := len(tiles)
	tileElements := make([][]Tile, size)
	for i := 0; i < size; i++ {
		tileElements[i] = make([]Tile, size)
		for j := 0; j < size; j++ {
			element := Tile{}
			if tiles[i][j] == 0 {
				element.Empty = true
			} else {
				element.DisplayValue = tiles[i][j]
				element.DomID = DomIDPrefix + strconv.Itoa(tiles[i][j])
			}
			tileElements[i][j] = element
		}
	}

	return &Board{
		Size:  size,
		Tiles: tileElements,
	}
}
