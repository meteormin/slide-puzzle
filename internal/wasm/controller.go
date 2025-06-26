package wasm

import (
	"syscall/js"
)

const (
	BoardID = "puzzle-board"
)

type Controller struct {
	document     js.Value
	boardSection js.Value
	board        *Board
}

// Init initializes the puzzle board with the given tiles.
func (c *Controller) Init(tiles [][]int) {
	for _, tile := range c.board.Tiles {
		for _, t := range tile {
			tileElement := c.document.Call("createElement", "div")
			tileElement.Get("classList").Call("add", "tile")
			if t.Empty {
				tileElement.Get("classList").Call("add", "empty")
			} else {
				tileElement.Set("textContent", t.DisplayValue)
			}
			c.boardSection.Call("appendChild", tileElement)
		}
	}
}

func NewController(titles [][]int) *Controller {

	document := js.Global().Get("document")
	boardSection := document.Call("getElementById", BoardID)

	c := &Controller{
		document:     document,
		boardSection: boardSection,
		board:        NewBoard(titles),
	}

	c.Init(titles)

	return c
}
